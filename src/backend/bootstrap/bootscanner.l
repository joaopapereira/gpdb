%{
/*-------------------------------------------------------------------------
 *
 * bootscanner.l
 *	  a lexical scanner for the bootstrap parser
 *
 * Portions Copyright (c) 1996-2009, PostgreSQL Global Development Group
 * Portions Copyright (c) 1994, Regents of the University of California
 *
 *
 * IDENTIFICATION
 *	  $PostgreSQL: pgsql/src/backend/bootstrap/bootscanner.l,v 1.48 2009/01/01 17:23:36 momjian Exp $
 *
 *-------------------------------------------------------------------------
 */
#include "postgres.h"

#include "access/attnum.h"
#include "access/htup.h"
#include "access/itup.h"
#include "access/skey.h"
#include "access/tupdesc.h"
#include "bootstrap/bootstrap.h"
#include "catalog/pg_am.h"
#include "catalog/pg_attribute.h"
#include "catalog/pg_class.h"
#include "nodes/nodes.h"
#include "nodes/parsenodes.h"
#include "nodes/pg_list.h"
#include "nodes/primnodes.h"
#include "parser/scansup.h"
#include "rewrite/prs2lock.h"
#include "storage/block.h"
#include "storage/fd.h"
#include "storage/itemptr.h"
#include "storage/off.h"
#include "utils/rel.h"

#define unify_version(a,b,c) ((a<<16)+(b<<8)+c)

#if unify_version(YY_FLEX_MAJOR_VERSION,YY_FLEX_MINOR_VERSION,YY_FLEX_SUBMINOR_VERSION) < unify_version(2,5,35)
/* These routines aren't used, but are defined to supress warnings from gcc */
int boot_yyget_lineno  (void);
FILE *boot_yyget_in  (void);
FILE *boot_yyget_out  (void);
int boot_yyget_leng  (void);
char *boot_yyget_text  (void);
void boot_yyset_lineno (int  line_number );
void boot_yyset_in (FILE *  in_str );
void boot_yyset_out (FILE *  out_str );
int boot_yyget_debug  (void);
void boot_yyset_debug (int  bdebug );
int boot_yylex_destroy  (void);
#endif





/* Not needed now that this file is compiled as part of bootparse. */
/* #include "bootparse.h" */


/* Avoid exit() on fatal scanner errors (a bit ugly -- see yy_fatal_error) */
#undef fprintf
#define fprintf(file, fmt, msg)  fprintf_to_ereport(fmt, msg)

static void
fprintf_to_ereport(const char *fmt, const char *msg)
{
	ereport(ERROR, (errmsg_internal("%s", msg)));
}


static int	yyline = 1;			/* line number for error reporting */

%}

%option 8bit
%option never-interactive
%option nodefault
%option noinput
%option nounput
%option noyywrap
%option prefix="boot_yy"


D		[0-9]
oct		\\{D}{D}{D}
Exp		[Ee][-+]?{D}+
id		([A-Za-z0-9_]|{oct}|\-)+
sid		\"([^\"])*\"
arrayid [A-Za-z0-9_]+\[{D}*\]

%%

open			{ return(OPEN); }

close			{ return(XCLOSE); }

create			{ return(XCREATE); }

OID				{ return(OBJ_ID); }
bootstrap		{ return(XBOOTSTRAP); }
"shared_relation"	{ return(XSHARED_RELATION); }
"without_oids"	{ return(XWITHOUT_OIDS); }
_null_			{ return(NULLVAL); }

insert			{ return(INSERT_TUPLE); }

","				{ return(COMMA); }
"="				{ return(EQUALS); }
"("				{ return(LPAREN); }
")"				{ return(RPAREN); }

[\n]			{ yyline++; }
[\t]			;
" "				;

^\#[^\n]* ; /* drop everything after "#" for comments */


"declare"		{ return(XDECLARE); }
"build"			{ return(XBUILD); }
"indices"		{ return(INDICES); }
"unique"		{ return(UNIQUE); }
"index"			{ return(INDEX); }
"on"			{ return(ON); }
"using"			{ return(USING); }
"toast"			{ return(XTOAST); }

{arrayid}		{
					yylval.ival = EnterString(MapArrayTypeName((char*)yytext));
					return(ID);
				}
{id}			{
					char   *newid = scanstr((char*)yytext);
					yylval.ival = EnterString(newid);
					pfree(newid);
					return(ID);
				}
{sid}			{
					char   *newid;
					yytext[strlen(yytext)-1] = '\0'; /* strip off quotes */
					newid = scanstr((char*)yytext+1);
					yylval.ival = EnterString(newid);
					pfree(newid);
					yytext[strlen(yytext)] = '"'; /* restore quotes */
					return(ID);
				}

(-)?{D}+"."{D}*({Exp})? |
(-)?{D}*"."{D}+({Exp})? |
(-)?{D}+{Exp}			{
							yylval.ival = EnterString((char*)yytext);
							return(CONST_P);
						}

.				{
					elog(ERROR, "syntax error at line %d: unexpected character \"%s\"", yyline, yytext);
				}



%%

void
yyerror(const char *message)
{
	elog(ERROR, "%s at line %d", message, yyline);
}
