/*-------------------------------------------------------------------------
 *
 * cdbexplain.h
 *    Functions supporting the Greenplum EXPLAIN ANALYZE command
 *
 * Portions Copyright (c) 2006-2008, Greenplum inc
 * Portions Copyright (c) 2012-Present Pivotal Software, Inc.
 *
 *
 * IDENTIFICATION
 *	    src/include/cdb/cdbexplain.h
 *
 *-------------------------------------------------------------------------
 */
#ifndef CDBEXPLAIN_H
#define CDBEXPLAIN_H

#include "executor/instrument.h"        /* instr_time */
#include "cdb/cdbpublic.h"              /* CdbExplain_Agg */

struct CdbDispatchResults;              /* #include "cdb/cdbdispatchresult.h" */
struct EState;                          /* #include "nodes/execnodes.h" */
struct PlanState;                       /* #include "nodes/execnodes.h" */
struct QueryDesc;                       /* #include "executor/execdesc.h" */
struct StringInfoData;                  /* #include "lib/stringinfo.h" */

struct GPExplain_ShowStatCtx;          /* private, in "cdb/cdbexplain.c" */
struct PlannedStmt;

static inline void
cdbexplain_agg_init0(CdbExplain_Agg *agg)
{
    agg->vmax = 0;
    agg->vsum = 0;
    agg->vcnt = 0;
    agg->imax = 0;
}

static inline bool
cdbexplain_agg_upd(CdbExplain_Agg *agg, double v, int id)
{
    if (v > 0)
    {
        agg->vsum += v;
        agg->vcnt++;

        if (v > agg->vmax ||
            agg->vcnt == 0)
        {
            agg->vmax = v;
            agg->imax = id;
            return true;
        }
    }
    return false;
}

static inline double
cdbexplain_agg_avg(CdbExplain_Agg *agg)
{
    return (agg->vcnt > 0) ? agg->vsum / agg->vcnt
                           : 0;
}


/*
 * gpexplain_localExecStats
 *    Called by qDisp to build NodeSummary and SliceSummary blocks
 *    containing EXPLAIN ANALYZE statistics for a root slice that
 *    has been executed locally in the qDisp process.  Attaches these
 *    structures to the PlanState nodes' Instrumentation objects for
 *    later use by gpexplain_showExecStats().
 *
 * 'planstate' is the top PlanState node of the slice.
 * 'showstatctx' is a GPExplain_ShowStatCtx object which was created by
 *      calling gpexplain_showExecStatsBegin().
 */
void
gpexplain_localExecStats(struct PlanState *planstate,
                         struct GPExplain_ShowStatCtx *showstatctx);

/*
 * gpexplain_sendExecStats
 *    Called by qExec process to send EXPLAIN ANALYZE statistics to qDisp.
 *    On the qDisp, libpq will attach this data to the PGresult object.
 */
void
gpexplain_sendExecStats(struct QueryDesc *queryDesc);

/*
 * gpexplain_recvExecStats
 *    Called by qDisp to transfer a slice's EXPLAIN ANALYZE statistics
 *    from the CdbDispatchResults structures to the PlanState tree.
 *    Recursively does the same for slices that are descendants of the
 *    one specified.
 *
 * 'showstatctx' is a GPExplain_ShowStatCtx object which was created by
 *      calling gpexplain_showExecStatsBegin().
 */
void
gpexplain_recvExecStats(struct PlanState *planstate,
                        struct CdbDispatchResults *dispatchResults,
                        int sliceIndex,
                        struct GPExplain_ShowStatCtx *showstatctx);

/*
 * gpexplain_showExecStatsBegin
 *    Called by qDisp process to create a GPExplain_ShowStatCtx structure
 *    in which to accumulate overall statistics for a query.
 *
 * 'explaincxt' is a MemoryContext from which to allocate the ShowStatCtx as
 *      well as any needed buffers and the like.  The explaincxt ptr is saved
 *      in the ShowStatCtx.  The caller is expected to reset or destroy the
 *      explaincxt not too long after calling gpexplain_showExecStatsEnd(); so
 *      we don't bother to pfree() memory that we allocate from this context.
 * 'querystarttime' is the timestamp of the start of the query, in a
 *      platform-dependent format.
 */
struct GPExplain_ShowStatCtx *
gpexplain_showExecStatsBegin(struct QueryDesc *queryDesc,
                             instr_time querystarttime);



#endif   /* CDBEXPLAIN_H */
