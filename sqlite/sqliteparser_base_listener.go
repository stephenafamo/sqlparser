// Code generated from ./sqlite/SQLiteParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SQLiteParser

import "github.com/antlr4-go/antlr/v4"

// BaseSQLiteParserListener is a complete listener for a parse tree produced by SQLiteParser.
type BaseSQLiteParserListener struct{}

var _ SQLiteParserListener = &BaseSQLiteParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSQLiteParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSQLiteParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSQLiteParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSQLiteParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseSQLiteParserListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseSQLiteParserListener) ExitParse(ctx *ParseContext) {}

// EnterSql_stmt_list is called when production sql_stmt_list is entered.
func (s *BaseSQLiteParserListener) EnterSql_stmt_list(ctx *Sql_stmt_listContext) {}

// ExitSql_stmt_list is called when production sql_stmt_list is exited.
func (s *BaseSQLiteParserListener) ExitSql_stmt_list(ctx *Sql_stmt_listContext) {}

// EnterSql_stmt is called when production sql_stmt is entered.
func (s *BaseSQLiteParserListener) EnterSql_stmt(ctx *Sql_stmtContext) {}

// ExitSql_stmt is called when production sql_stmt is exited.
func (s *BaseSQLiteParserListener) ExitSql_stmt(ctx *Sql_stmtContext) {}

// EnterAlter_table_stmt is called when production alter_table_stmt is entered.
func (s *BaseSQLiteParserListener) EnterAlter_table_stmt(ctx *Alter_table_stmtContext) {}

// ExitAlter_table_stmt is called when production alter_table_stmt is exited.
func (s *BaseSQLiteParserListener) ExitAlter_table_stmt(ctx *Alter_table_stmtContext) {}

// EnterAnalyze_stmt is called when production analyze_stmt is entered.
func (s *BaseSQLiteParserListener) EnterAnalyze_stmt(ctx *Analyze_stmtContext) {}

// ExitAnalyze_stmt is called when production analyze_stmt is exited.
func (s *BaseSQLiteParserListener) ExitAnalyze_stmt(ctx *Analyze_stmtContext) {}

// EnterAttach_stmt is called when production attach_stmt is entered.
func (s *BaseSQLiteParserListener) EnterAttach_stmt(ctx *Attach_stmtContext) {}

// ExitAttach_stmt is called when production attach_stmt is exited.
func (s *BaseSQLiteParserListener) ExitAttach_stmt(ctx *Attach_stmtContext) {}

// EnterBegin_stmt is called when production begin_stmt is entered.
func (s *BaseSQLiteParserListener) EnterBegin_stmt(ctx *Begin_stmtContext) {}

// ExitBegin_stmt is called when production begin_stmt is exited.
func (s *BaseSQLiteParserListener) ExitBegin_stmt(ctx *Begin_stmtContext) {}

// EnterCommit_stmt is called when production commit_stmt is entered.
func (s *BaseSQLiteParserListener) EnterCommit_stmt(ctx *Commit_stmtContext) {}

// ExitCommit_stmt is called when production commit_stmt is exited.
func (s *BaseSQLiteParserListener) ExitCommit_stmt(ctx *Commit_stmtContext) {}

// EnterRollback_stmt is called when production rollback_stmt is entered.
func (s *BaseSQLiteParserListener) EnterRollback_stmt(ctx *Rollback_stmtContext) {}

// ExitRollback_stmt is called when production rollback_stmt is exited.
func (s *BaseSQLiteParserListener) ExitRollback_stmt(ctx *Rollback_stmtContext) {}

// EnterSavepoint_stmt is called when production savepoint_stmt is entered.
func (s *BaseSQLiteParserListener) EnterSavepoint_stmt(ctx *Savepoint_stmtContext) {}

// ExitSavepoint_stmt is called when production savepoint_stmt is exited.
func (s *BaseSQLiteParserListener) ExitSavepoint_stmt(ctx *Savepoint_stmtContext) {}

// EnterRelease_stmt is called when production release_stmt is entered.
func (s *BaseSQLiteParserListener) EnterRelease_stmt(ctx *Release_stmtContext) {}

// ExitRelease_stmt is called when production release_stmt is exited.
func (s *BaseSQLiteParserListener) ExitRelease_stmt(ctx *Release_stmtContext) {}

// EnterCreate_index_stmt is called when production create_index_stmt is entered.
func (s *BaseSQLiteParserListener) EnterCreate_index_stmt(ctx *Create_index_stmtContext) {}

// ExitCreate_index_stmt is called when production create_index_stmt is exited.
func (s *BaseSQLiteParserListener) ExitCreate_index_stmt(ctx *Create_index_stmtContext) {}

// EnterIndexed_column is called when production indexed_column is entered.
func (s *BaseSQLiteParserListener) EnterIndexed_column(ctx *Indexed_columnContext) {}

// ExitIndexed_column is called when production indexed_column is exited.
func (s *BaseSQLiteParserListener) ExitIndexed_column(ctx *Indexed_columnContext) {}

// EnterCreate_table_stmt is called when production create_table_stmt is entered.
func (s *BaseSQLiteParserListener) EnterCreate_table_stmt(ctx *Create_table_stmtContext) {}

// ExitCreate_table_stmt is called when production create_table_stmt is exited.
func (s *BaseSQLiteParserListener) ExitCreate_table_stmt(ctx *Create_table_stmtContext) {}

// EnterColumn_def is called when production column_def is entered.
func (s *BaseSQLiteParserListener) EnterColumn_def(ctx *Column_defContext) {}

// ExitColumn_def is called when production column_def is exited.
func (s *BaseSQLiteParserListener) ExitColumn_def(ctx *Column_defContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BaseSQLiteParserListener) EnterType_name(ctx *Type_nameContext) {}

// ExitType_name is called when production type_name is exited.
func (s *BaseSQLiteParserListener) ExitType_name(ctx *Type_nameContext) {}

// EnterColumn_constraint is called when production column_constraint is entered.
func (s *BaseSQLiteParserListener) EnterColumn_constraint(ctx *Column_constraintContext) {}

// ExitColumn_constraint is called when production column_constraint is exited.
func (s *BaseSQLiteParserListener) ExitColumn_constraint(ctx *Column_constraintContext) {}

// EnterSigned_number is called when production signed_number is entered.
func (s *BaseSQLiteParserListener) EnterSigned_number(ctx *Signed_numberContext) {}

// ExitSigned_number is called when production signed_number is exited.
func (s *BaseSQLiteParserListener) ExitSigned_number(ctx *Signed_numberContext) {}

// EnterTable_constraint is called when production table_constraint is entered.
func (s *BaseSQLiteParserListener) EnterTable_constraint(ctx *Table_constraintContext) {}

// ExitTable_constraint is called when production table_constraint is exited.
func (s *BaseSQLiteParserListener) ExitTable_constraint(ctx *Table_constraintContext) {}

// EnterForeign_key_clause is called when production foreign_key_clause is entered.
func (s *BaseSQLiteParserListener) EnterForeign_key_clause(ctx *Foreign_key_clauseContext) {}

// ExitForeign_key_clause is called when production foreign_key_clause is exited.
func (s *BaseSQLiteParserListener) ExitForeign_key_clause(ctx *Foreign_key_clauseContext) {}

// EnterConflict_clause is called when production conflict_clause is entered.
func (s *BaseSQLiteParserListener) EnterConflict_clause(ctx *Conflict_clauseContext) {}

// ExitConflict_clause is called when production conflict_clause is exited.
func (s *BaseSQLiteParserListener) ExitConflict_clause(ctx *Conflict_clauseContext) {}

// EnterCreate_trigger_stmt is called when production create_trigger_stmt is entered.
func (s *BaseSQLiteParserListener) EnterCreate_trigger_stmt(ctx *Create_trigger_stmtContext) {}

// ExitCreate_trigger_stmt is called when production create_trigger_stmt is exited.
func (s *BaseSQLiteParserListener) ExitCreate_trigger_stmt(ctx *Create_trigger_stmtContext) {}

// EnterCreate_view_stmt is called when production create_view_stmt is entered.
func (s *BaseSQLiteParserListener) EnterCreate_view_stmt(ctx *Create_view_stmtContext) {}

// ExitCreate_view_stmt is called when production create_view_stmt is exited.
func (s *BaseSQLiteParserListener) ExitCreate_view_stmt(ctx *Create_view_stmtContext) {}

// EnterCreate_virtual_table_stmt is called when production create_virtual_table_stmt is entered.
func (s *BaseSQLiteParserListener) EnterCreate_virtual_table_stmt(ctx *Create_virtual_table_stmtContext) {
}

// ExitCreate_virtual_table_stmt is called when production create_virtual_table_stmt is exited.
func (s *BaseSQLiteParserListener) ExitCreate_virtual_table_stmt(ctx *Create_virtual_table_stmtContext) {
}

// EnterDelete_stmt is called when production delete_stmt is entered.
func (s *BaseSQLiteParserListener) EnterDelete_stmt(ctx *Delete_stmtContext) {}

// ExitDelete_stmt is called when production delete_stmt is exited.
func (s *BaseSQLiteParserListener) ExitDelete_stmt(ctx *Delete_stmtContext) {}

// EnterDetach_stmt is called when production detach_stmt is entered.
func (s *BaseSQLiteParserListener) EnterDetach_stmt(ctx *Detach_stmtContext) {}

// ExitDetach_stmt is called when production detach_stmt is exited.
func (s *BaseSQLiteParserListener) ExitDetach_stmt(ctx *Detach_stmtContext) {}

// EnterDrop_stmt is called when production drop_stmt is entered.
func (s *BaseSQLiteParserListener) EnterDrop_stmt(ctx *Drop_stmtContext) {}

// ExitDrop_stmt is called when production drop_stmt is exited.
func (s *BaseSQLiteParserListener) ExitDrop_stmt(ctx *Drop_stmtContext) {}

// EnterExpr_case is called when production expr_case is entered.
func (s *BaseSQLiteParserListener) EnterExpr_case(ctx *Expr_caseContext) {}

// ExitExpr_case is called when production expr_case is exited.
func (s *BaseSQLiteParserListener) ExitExpr_case(ctx *Expr_caseContext) {}

// EnterExpr_arithmetic is called when production expr_arithmetic is entered.
func (s *BaseSQLiteParserListener) EnterExpr_arithmetic(ctx *Expr_arithmeticContext) {}

// ExitExpr_arithmetic is called when production expr_arithmetic is exited.
func (s *BaseSQLiteParserListener) ExitExpr_arithmetic(ctx *Expr_arithmeticContext) {}

// EnterExpr_aggregate_func is called when production expr_aggregate_func is entered.
func (s *BaseSQLiteParserListener) EnterExpr_aggregate_func(ctx *Expr_aggregate_funcContext) {}

// ExitExpr_aggregate_func is called when production expr_aggregate_func is exited.
func (s *BaseSQLiteParserListener) ExitExpr_aggregate_func(ctx *Expr_aggregate_funcContext) {}

// EnterExpr_json_extract_string is called when production expr_json_extract_string is entered.
func (s *BaseSQLiteParserListener) EnterExpr_json_extract_string(ctx *Expr_json_extract_stringContext) {
}

// ExitExpr_json_extract_string is called when production expr_json_extract_string is exited.
func (s *BaseSQLiteParserListener) ExitExpr_json_extract_string(ctx *Expr_json_extract_stringContext) {
}

// EnterExpr_raise is called when production expr_raise is entered.
func (s *BaseSQLiteParserListener) EnterExpr_raise(ctx *Expr_raiseContext) {}

// ExitExpr_raise is called when production expr_raise is exited.
func (s *BaseSQLiteParserListener) ExitExpr_raise(ctx *Expr_raiseContext) {}

// EnterExpr_bool is called when production expr_bool is entered.
func (s *BaseSQLiteParserListener) EnterExpr_bool(ctx *Expr_boolContext) {}

// ExitExpr_bool is called when production expr_bool is exited.
func (s *BaseSQLiteParserListener) ExitExpr_bool(ctx *Expr_boolContext) {}

// EnterExpr_is is called when production expr_is is entered.
func (s *BaseSQLiteParserListener) EnterExpr_is(ctx *Expr_isContext) {}

// ExitExpr_is is called when production expr_is is exited.
func (s *BaseSQLiteParserListener) ExitExpr_is(ctx *Expr_isContext) {}

// EnterExpr_concat is called when production expr_concat is entered.
func (s *BaseSQLiteParserListener) EnterExpr_concat(ctx *Expr_concatContext) {}

// ExitExpr_concat is called when production expr_concat is exited.
func (s *BaseSQLiteParserListener) ExitExpr_concat(ctx *Expr_concatContext) {}

// EnterExpr_list is called when production expr_list is entered.
func (s *BaseSQLiteParserListener) EnterExpr_list(ctx *Expr_listContext) {}

// ExitExpr_list is called when production expr_list is exited.
func (s *BaseSQLiteParserListener) ExitExpr_list(ctx *Expr_listContext) {}

// EnterExpr_in is called when production expr_in is entered.
func (s *BaseSQLiteParserListener) EnterExpr_in(ctx *Expr_inContext) {}

// ExitExpr_in is called when production expr_in is exited.
func (s *BaseSQLiteParserListener) ExitExpr_in(ctx *Expr_inContext) {}

// EnterExpr_collate is called when production expr_collate is entered.
func (s *BaseSQLiteParserListener) EnterExpr_collate(ctx *Expr_collateContext) {}

// ExitExpr_collate is called when production expr_collate is exited.
func (s *BaseSQLiteParserListener) ExitExpr_collate(ctx *Expr_collateContext) {}

// EnterExpr_modulo is called when production expr_modulo is entered.
func (s *BaseSQLiteParserListener) EnterExpr_modulo(ctx *Expr_moduloContext) {}

// ExitExpr_modulo is called when production expr_modulo is exited.
func (s *BaseSQLiteParserListener) ExitExpr_modulo(ctx *Expr_moduloContext) {}

// EnterExpr_qualified_column_name is called when production expr_qualified_column_name is entered.
func (s *BaseSQLiteParserListener) EnterExpr_qualified_column_name(ctx *Expr_qualified_column_nameContext) {
}

// ExitExpr_qualified_column_name is called when production expr_qualified_column_name is exited.
func (s *BaseSQLiteParserListener) ExitExpr_qualified_column_name(ctx *Expr_qualified_column_nameContext) {
}

// EnterExpr_match is called when production expr_match is entered.
func (s *BaseSQLiteParserListener) EnterExpr_match(ctx *Expr_matchContext) {}

// ExitExpr_match is called when production expr_match is exited.
func (s *BaseSQLiteParserListener) ExitExpr_match(ctx *Expr_matchContext) {}

// EnterExpr_like is called when production expr_like is entered.
func (s *BaseSQLiteParserListener) EnterExpr_like(ctx *Expr_likeContext) {}

// ExitExpr_like is called when production expr_like is exited.
func (s *BaseSQLiteParserListener) ExitExpr_like(ctx *Expr_likeContext) {}

// EnterExpr_null_comp is called when production expr_null_comp is entered.
func (s *BaseSQLiteParserListener) EnterExpr_null_comp(ctx *Expr_null_compContext) {}

// ExitExpr_null_comp is called when production expr_null_comp is exited.
func (s *BaseSQLiteParserListener) ExitExpr_null_comp(ctx *Expr_null_compContext) {}

// EnterExpr_json_extract_json is called when production expr_json_extract_json is entered.
func (s *BaseSQLiteParserListener) EnterExpr_json_extract_json(ctx *Expr_json_extract_jsonContext) {}

// ExitExpr_json_extract_json is called when production expr_json_extract_json is exited.
func (s *BaseSQLiteParserListener) ExitExpr_json_extract_json(ctx *Expr_json_extract_jsonContext) {}

// EnterExpr_window_func is called when production expr_window_func is entered.
func (s *BaseSQLiteParserListener) EnterExpr_window_func(ctx *Expr_window_funcContext) {}

// ExitExpr_window_func is called when production expr_window_func is exited.
func (s *BaseSQLiteParserListener) ExitExpr_window_func(ctx *Expr_window_funcContext) {}

// EnterExpr_exists_select is called when production expr_exists_select is entered.
func (s *BaseSQLiteParserListener) EnterExpr_exists_select(ctx *Expr_exists_selectContext) {}

// ExitExpr_exists_select is called when production expr_exists_select is exited.
func (s *BaseSQLiteParserListener) ExitExpr_exists_select(ctx *Expr_exists_selectContext) {}

// EnterExpr_comparison is called when production expr_comparison is entered.
func (s *BaseSQLiteParserListener) EnterExpr_comparison(ctx *Expr_comparisonContext) {}

// ExitExpr_comparison is called when production expr_comparison is exited.
func (s *BaseSQLiteParserListener) ExitExpr_comparison(ctx *Expr_comparisonContext) {}

// EnterExpr_literal is called when production expr_literal is entered.
func (s *BaseSQLiteParserListener) EnterExpr_literal(ctx *Expr_literalContext) {}

// ExitExpr_literal is called when production expr_literal is exited.
func (s *BaseSQLiteParserListener) ExitExpr_literal(ctx *Expr_literalContext) {}

// EnterExpr_cast is called when production expr_cast is entered.
func (s *BaseSQLiteParserListener) EnterExpr_cast(ctx *Expr_castContext) {}

// ExitExpr_cast is called when production expr_cast is exited.
func (s *BaseSQLiteParserListener) ExitExpr_cast(ctx *Expr_castContext) {}

// EnterExpr_string_op is called when production expr_string_op is entered.
func (s *BaseSQLiteParserListener) EnterExpr_string_op(ctx *Expr_string_opContext) {}

// ExitExpr_string_op is called when production expr_string_op is exited.
func (s *BaseSQLiteParserListener) ExitExpr_string_op(ctx *Expr_string_opContext) {}

// EnterExpr_between is called when production expr_between is entered.
func (s *BaseSQLiteParserListener) EnterExpr_between(ctx *Expr_betweenContext) {}

// ExitExpr_between is called when production expr_between is exited.
func (s *BaseSQLiteParserListener) ExitExpr_between(ctx *Expr_betweenContext) {}

// EnterExpr_bitwise is called when production expr_bitwise is entered.
func (s *BaseSQLiteParserListener) EnterExpr_bitwise(ctx *Expr_bitwiseContext) {}

// ExitExpr_bitwise is called when production expr_bitwise is exited.
func (s *BaseSQLiteParserListener) ExitExpr_bitwise(ctx *Expr_bitwiseContext) {}

// EnterExpr_simple_func is called when production expr_simple_func is entered.
func (s *BaseSQLiteParserListener) EnterExpr_simple_func(ctx *Expr_simple_funcContext) {}

// ExitExpr_simple_func is called when production expr_simple_func is exited.
func (s *BaseSQLiteParserListener) ExitExpr_simple_func(ctx *Expr_simple_funcContext) {}

// EnterExpr_unary is called when production expr_unary is entered.
func (s *BaseSQLiteParserListener) EnterExpr_unary(ctx *Expr_unaryContext) {}

// ExitExpr_unary is called when production expr_unary is exited.
func (s *BaseSQLiteParserListener) ExitExpr_unary(ctx *Expr_unaryContext) {}

// EnterExpr_bind is called when production expr_bind is entered.
func (s *BaseSQLiteParserListener) EnterExpr_bind(ctx *Expr_bindContext) {}

// ExitExpr_bind is called when production expr_bind is exited.
func (s *BaseSQLiteParserListener) ExitExpr_bind(ctx *Expr_bindContext) {}

// EnterRaise_function is called when production raise_function is entered.
func (s *BaseSQLiteParserListener) EnterRaise_function(ctx *Raise_functionContext) {}

// ExitRaise_function is called when production raise_function is exited.
func (s *BaseSQLiteParserListener) ExitRaise_function(ctx *Raise_functionContext) {}

// EnterLiteral_value is called when production literal_value is entered.
func (s *BaseSQLiteParserListener) EnterLiteral_value(ctx *Literal_valueContext) {}

// ExitLiteral_value is called when production literal_value is exited.
func (s *BaseSQLiteParserListener) ExitLiteral_value(ctx *Literal_valueContext) {}

// EnterValue_row is called when production value_row is entered.
func (s *BaseSQLiteParserListener) EnterValue_row(ctx *Value_rowContext) {}

// ExitValue_row is called when production value_row is exited.
func (s *BaseSQLiteParserListener) ExitValue_row(ctx *Value_rowContext) {}

// EnterValues_clause is called when production values_clause is entered.
func (s *BaseSQLiteParserListener) EnterValues_clause(ctx *Values_clauseContext) {}

// ExitValues_clause is called when production values_clause is exited.
func (s *BaseSQLiteParserListener) ExitValues_clause(ctx *Values_clauseContext) {}

// EnterInsert_stmt is called when production insert_stmt is entered.
func (s *BaseSQLiteParserListener) EnterInsert_stmt(ctx *Insert_stmtContext) {}

// ExitInsert_stmt is called when production insert_stmt is exited.
func (s *BaseSQLiteParserListener) ExitInsert_stmt(ctx *Insert_stmtContext) {}

// EnterReturning_clause is called when production returning_clause is entered.
func (s *BaseSQLiteParserListener) EnterReturning_clause(ctx *Returning_clauseContext) {}

// ExitReturning_clause is called when production returning_clause is exited.
func (s *BaseSQLiteParserListener) ExitReturning_clause(ctx *Returning_clauseContext) {}

// EnterUpsert_clause is called when production upsert_clause is entered.
func (s *BaseSQLiteParserListener) EnterUpsert_clause(ctx *Upsert_clauseContext) {}

// ExitUpsert_clause is called when production upsert_clause is exited.
func (s *BaseSQLiteParserListener) ExitUpsert_clause(ctx *Upsert_clauseContext) {}

// EnterPragma_stmt is called when production pragma_stmt is entered.
func (s *BaseSQLiteParserListener) EnterPragma_stmt(ctx *Pragma_stmtContext) {}

// ExitPragma_stmt is called when production pragma_stmt is exited.
func (s *BaseSQLiteParserListener) ExitPragma_stmt(ctx *Pragma_stmtContext) {}

// EnterPragma_value is called when production pragma_value is entered.
func (s *BaseSQLiteParserListener) EnterPragma_value(ctx *Pragma_valueContext) {}

// ExitPragma_value is called when production pragma_value is exited.
func (s *BaseSQLiteParserListener) ExitPragma_value(ctx *Pragma_valueContext) {}

// EnterReindex_stmt is called when production reindex_stmt is entered.
func (s *BaseSQLiteParserListener) EnterReindex_stmt(ctx *Reindex_stmtContext) {}

// ExitReindex_stmt is called when production reindex_stmt is exited.
func (s *BaseSQLiteParserListener) ExitReindex_stmt(ctx *Reindex_stmtContext) {}

// EnterSelect_stmt is called when production select_stmt is entered.
func (s *BaseSQLiteParserListener) EnterSelect_stmt(ctx *Select_stmtContext) {}

// ExitSelect_stmt is called when production select_stmt is exited.
func (s *BaseSQLiteParserListener) ExitSelect_stmt(ctx *Select_stmtContext) {}

// EnterSelect_core is called when production select_core is entered.
func (s *BaseSQLiteParserListener) EnterSelect_core(ctx *Select_coreContext) {}

// ExitSelect_core is called when production select_core is exited.
func (s *BaseSQLiteParserListener) ExitSelect_core(ctx *Select_coreContext) {}

// EnterResult_column is called when production result_column is entered.
func (s *BaseSQLiteParserListener) EnterResult_column(ctx *Result_columnContext) {}

// ExitResult_column is called when production result_column is exited.
func (s *BaseSQLiteParserListener) ExitResult_column(ctx *Result_columnContext) {}

// EnterFrom_item is called when production from_item is entered.
func (s *BaseSQLiteParserListener) EnterFrom_item(ctx *From_itemContext) {}

// ExitFrom_item is called when production from_item is exited.
func (s *BaseSQLiteParserListener) ExitFrom_item(ctx *From_itemContext) {}

// EnterJoin_operator is called when production join_operator is entered.
func (s *BaseSQLiteParserListener) EnterJoin_operator(ctx *Join_operatorContext) {}

// ExitJoin_operator is called when production join_operator is exited.
func (s *BaseSQLiteParserListener) ExitJoin_operator(ctx *Join_operatorContext) {}

// EnterJoin_constraint is called when production join_constraint is entered.
func (s *BaseSQLiteParserListener) EnterJoin_constraint(ctx *Join_constraintContext) {}

// ExitJoin_constraint is called when production join_constraint is exited.
func (s *BaseSQLiteParserListener) ExitJoin_constraint(ctx *Join_constraintContext) {}

// EnterTable_or_subquery is called when production table_or_subquery is entered.
func (s *BaseSQLiteParserListener) EnterTable_or_subquery(ctx *Table_or_subqueryContext) {}

// ExitTable_or_subquery is called when production table_or_subquery is exited.
func (s *BaseSQLiteParserListener) ExitTable_or_subquery(ctx *Table_or_subqueryContext) {}

// EnterCompound_select is called when production compound_select is entered.
func (s *BaseSQLiteParserListener) EnterCompound_select(ctx *Compound_selectContext) {}

// ExitCompound_select is called when production compound_select is exited.
func (s *BaseSQLiteParserListener) ExitCompound_select(ctx *Compound_selectContext) {}

// EnterCompound_operator is called when production compound_operator is entered.
func (s *BaseSQLiteParserListener) EnterCompound_operator(ctx *Compound_operatorContext) {}

// ExitCompound_operator is called when production compound_operator is exited.
func (s *BaseSQLiteParserListener) ExitCompound_operator(ctx *Compound_operatorContext) {}

// EnterUpdate_stmt is called when production update_stmt is entered.
func (s *BaseSQLiteParserListener) EnterUpdate_stmt(ctx *Update_stmtContext) {}

// ExitUpdate_stmt is called when production update_stmt is exited.
func (s *BaseSQLiteParserListener) ExitUpdate_stmt(ctx *Update_stmtContext) {}

// EnterColumn_name_list is called when production column_name_list is entered.
func (s *BaseSQLiteParserListener) EnterColumn_name_list(ctx *Column_name_listContext) {}

// ExitColumn_name_list is called when production column_name_list is exited.
func (s *BaseSQLiteParserListener) ExitColumn_name_list(ctx *Column_name_listContext) {}

// EnterQualified_table_name is called when production qualified_table_name is entered.
func (s *BaseSQLiteParserListener) EnterQualified_table_name(ctx *Qualified_table_nameContext) {}

// ExitQualified_table_name is called when production qualified_table_name is exited.
func (s *BaseSQLiteParserListener) ExitQualified_table_name(ctx *Qualified_table_nameContext) {}

// EnterVacuum_stmt is called when production vacuum_stmt is entered.
func (s *BaseSQLiteParserListener) EnterVacuum_stmt(ctx *Vacuum_stmtContext) {}

// ExitVacuum_stmt is called when production vacuum_stmt is exited.
func (s *BaseSQLiteParserListener) ExitVacuum_stmt(ctx *Vacuum_stmtContext) {}

// EnterFilter_clause is called when production filter_clause is entered.
func (s *BaseSQLiteParserListener) EnterFilter_clause(ctx *Filter_clauseContext) {}

// ExitFilter_clause is called when production filter_clause is exited.
func (s *BaseSQLiteParserListener) ExitFilter_clause(ctx *Filter_clauseContext) {}

// EnterWindow_defn is called when production window_defn is entered.
func (s *BaseSQLiteParserListener) EnterWindow_defn(ctx *Window_defnContext) {}

// ExitWindow_defn is called when production window_defn is exited.
func (s *BaseSQLiteParserListener) ExitWindow_defn(ctx *Window_defnContext) {}

// EnterOver_clause is called when production over_clause is entered.
func (s *BaseSQLiteParserListener) EnterOver_clause(ctx *Over_clauseContext) {}

// ExitOver_clause is called when production over_clause is exited.
func (s *BaseSQLiteParserListener) ExitOver_clause(ctx *Over_clauseContext) {}

// EnterFrame_spec is called when production frame_spec is entered.
func (s *BaseSQLiteParserListener) EnterFrame_spec(ctx *Frame_specContext) {}

// ExitFrame_spec is called when production frame_spec is exited.
func (s *BaseSQLiteParserListener) ExitFrame_spec(ctx *Frame_specContext) {}

// EnterFrame_clause is called when production frame_clause is entered.
func (s *BaseSQLiteParserListener) EnterFrame_clause(ctx *Frame_clauseContext) {}

// ExitFrame_clause is called when production frame_clause is exited.
func (s *BaseSQLiteParserListener) ExitFrame_clause(ctx *Frame_clauseContext) {}

// EnterWith_clause is called when production with_clause is entered.
func (s *BaseSQLiteParserListener) EnterWith_clause(ctx *With_clauseContext) {}

// ExitWith_clause is called when production with_clause is exited.
func (s *BaseSQLiteParserListener) ExitWith_clause(ctx *With_clauseContext) {}

// EnterCommon_table_expression is called when production common_table_expression is entered.
func (s *BaseSQLiteParserListener) EnterCommon_table_expression(ctx *Common_table_expressionContext) {
}

// ExitCommon_table_expression is called when production common_table_expression is exited.
func (s *BaseSQLiteParserListener) ExitCommon_table_expression(ctx *Common_table_expressionContext) {}

// EnterWhere_stmt is called when production where_stmt is entered.
func (s *BaseSQLiteParserListener) EnterWhere_stmt(ctx *Where_stmtContext) {}

// ExitWhere_stmt is called when production where_stmt is exited.
func (s *BaseSQLiteParserListener) ExitWhere_stmt(ctx *Where_stmtContext) {}

// EnterOrder_by_stmt is called when production order_by_stmt is entered.
func (s *BaseSQLiteParserListener) EnterOrder_by_stmt(ctx *Order_by_stmtContext) {}

// ExitOrder_by_stmt is called when production order_by_stmt is exited.
func (s *BaseSQLiteParserListener) ExitOrder_by_stmt(ctx *Order_by_stmtContext) {}

// EnterGroup_by_stmt is called when production group_by_stmt is entered.
func (s *BaseSQLiteParserListener) EnterGroup_by_stmt(ctx *Group_by_stmtContext) {}

// ExitGroup_by_stmt is called when production group_by_stmt is exited.
func (s *BaseSQLiteParserListener) ExitGroup_by_stmt(ctx *Group_by_stmtContext) {}

// EnterWindow_stmt is called when production window_stmt is entered.
func (s *BaseSQLiteParserListener) EnterWindow_stmt(ctx *Window_stmtContext) {}

// ExitWindow_stmt is called when production window_stmt is exited.
func (s *BaseSQLiteParserListener) ExitWindow_stmt(ctx *Window_stmtContext) {}

// EnterLimit_stmt is called when production limit_stmt is entered.
func (s *BaseSQLiteParserListener) EnterLimit_stmt(ctx *Limit_stmtContext) {}

// ExitLimit_stmt is called when production limit_stmt is exited.
func (s *BaseSQLiteParserListener) ExitLimit_stmt(ctx *Limit_stmtContext) {}

// EnterOrdering_term is called when production ordering_term is entered.
func (s *BaseSQLiteParserListener) EnterOrdering_term(ctx *Ordering_termContext) {}

// ExitOrdering_term is called when production ordering_term is exited.
func (s *BaseSQLiteParserListener) ExitOrdering_term(ctx *Ordering_termContext) {}

// EnterAsc_desc is called when production asc_desc is entered.
func (s *BaseSQLiteParserListener) EnterAsc_desc(ctx *Asc_descContext) {}

// ExitAsc_desc is called when production asc_desc is exited.
func (s *BaseSQLiteParserListener) ExitAsc_desc(ctx *Asc_descContext) {}

// EnterFrame_left is called when production frame_left is entered.
func (s *BaseSQLiteParserListener) EnterFrame_left(ctx *Frame_leftContext) {}

// ExitFrame_left is called when production frame_left is exited.
func (s *BaseSQLiteParserListener) ExitFrame_left(ctx *Frame_leftContext) {}

// EnterFrame_right is called when production frame_right is entered.
func (s *BaseSQLiteParserListener) EnterFrame_right(ctx *Frame_rightContext) {}

// ExitFrame_right is called when production frame_right is exited.
func (s *BaseSQLiteParserListener) ExitFrame_right(ctx *Frame_rightContext) {}

// EnterFrame_single is called when production frame_single is entered.
func (s *BaseSQLiteParserListener) EnterFrame_single(ctx *Frame_singleContext) {}

// ExitFrame_single is called when production frame_single is exited.
func (s *BaseSQLiteParserListener) ExitFrame_single(ctx *Frame_singleContext) {}

// EnterOffset is called when production offset is entered.
func (s *BaseSQLiteParserListener) EnterOffset(ctx *OffsetContext) {}

// ExitOffset is called when production offset is exited.
func (s *BaseSQLiteParserListener) ExitOffset(ctx *OffsetContext) {}

// EnterDefault_value is called when production default_value is entered.
func (s *BaseSQLiteParserListener) EnterDefault_value(ctx *Default_valueContext) {}

// ExitDefault_value is called when production default_value is exited.
func (s *BaseSQLiteParserListener) ExitDefault_value(ctx *Default_valueContext) {}

// EnterPartition_by is called when production partition_by is entered.
func (s *BaseSQLiteParserListener) EnterPartition_by(ctx *Partition_byContext) {}

// ExitPartition_by is called when production partition_by is exited.
func (s *BaseSQLiteParserListener) ExitPartition_by(ctx *Partition_byContext) {}

// EnterOrder_by_expr is called when production order_by_expr is entered.
func (s *BaseSQLiteParserListener) EnterOrder_by_expr(ctx *Order_by_exprContext) {}

// ExitOrder_by_expr is called when production order_by_expr is exited.
func (s *BaseSQLiteParserListener) ExitOrder_by_expr(ctx *Order_by_exprContext) {}

// EnterOrder_by_expr_asc_desc is called when production order_by_expr_asc_desc is entered.
func (s *BaseSQLiteParserListener) EnterOrder_by_expr_asc_desc(ctx *Order_by_expr_asc_descContext) {}

// ExitOrder_by_expr_asc_desc is called when production order_by_expr_asc_desc is exited.
func (s *BaseSQLiteParserListener) ExitOrder_by_expr_asc_desc(ctx *Order_by_expr_asc_descContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseSQLiteParserListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseSQLiteParserListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterError_message is called when production error_message is entered.
func (s *BaseSQLiteParserListener) EnterError_message(ctx *Error_messageContext) {}

// ExitError_message is called when production error_message is exited.
func (s *BaseSQLiteParserListener) ExitError_message(ctx *Error_messageContext) {}

// EnterModule_argument is called when production module_argument is entered.
func (s *BaseSQLiteParserListener) EnterModule_argument(ctx *Module_argumentContext) {}

// ExitModule_argument is called when production module_argument is exited.
func (s *BaseSQLiteParserListener) ExitModule_argument(ctx *Module_argumentContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseSQLiteParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseSQLiteParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterName is called when production name is entered.
func (s *BaseSQLiteParserListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseSQLiteParserListener) ExitName(ctx *NameContext) {}

// EnterFunction_name is called when production function_name is entered.
func (s *BaseSQLiteParserListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BaseSQLiteParserListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterSchema_name is called when production schema_name is entered.
func (s *BaseSQLiteParserListener) EnterSchema_name(ctx *Schema_nameContext) {}

// ExitSchema_name is called when production schema_name is exited.
func (s *BaseSQLiteParserListener) ExitSchema_name(ctx *Schema_nameContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseSQLiteParserListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseSQLiteParserListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterTable_or_index_name is called when production table_or_index_name is entered.
func (s *BaseSQLiteParserListener) EnterTable_or_index_name(ctx *Table_or_index_nameContext) {}

// ExitTable_or_index_name is called when production table_or_index_name is exited.
func (s *BaseSQLiteParserListener) ExitTable_or_index_name(ctx *Table_or_index_nameContext) {}

// EnterColumn_name is called when production column_name is entered.
func (s *BaseSQLiteParserListener) EnterColumn_name(ctx *Column_nameContext) {}

// ExitColumn_name is called when production column_name is exited.
func (s *BaseSQLiteParserListener) ExitColumn_name(ctx *Column_nameContext) {}

// EnterCollation_name is called when production collation_name is entered.
func (s *BaseSQLiteParserListener) EnterCollation_name(ctx *Collation_nameContext) {}

// ExitCollation_name is called when production collation_name is exited.
func (s *BaseSQLiteParserListener) ExitCollation_name(ctx *Collation_nameContext) {}

// EnterForeign_table is called when production foreign_table is entered.
func (s *BaseSQLiteParserListener) EnterForeign_table(ctx *Foreign_tableContext) {}

// ExitForeign_table is called when production foreign_table is exited.
func (s *BaseSQLiteParserListener) ExitForeign_table(ctx *Foreign_tableContext) {}

// EnterIndex_name is called when production index_name is entered.
func (s *BaseSQLiteParserListener) EnterIndex_name(ctx *Index_nameContext) {}

// ExitIndex_name is called when production index_name is exited.
func (s *BaseSQLiteParserListener) ExitIndex_name(ctx *Index_nameContext) {}

// EnterTrigger_name is called when production trigger_name is entered.
func (s *BaseSQLiteParserListener) EnterTrigger_name(ctx *Trigger_nameContext) {}

// ExitTrigger_name is called when production trigger_name is exited.
func (s *BaseSQLiteParserListener) ExitTrigger_name(ctx *Trigger_nameContext) {}

// EnterView_name is called when production view_name is entered.
func (s *BaseSQLiteParserListener) EnterView_name(ctx *View_nameContext) {}

// ExitView_name is called when production view_name is exited.
func (s *BaseSQLiteParserListener) ExitView_name(ctx *View_nameContext) {}

// EnterModule_name is called when production module_name is entered.
func (s *BaseSQLiteParserListener) EnterModule_name(ctx *Module_nameContext) {}

// ExitModule_name is called when production module_name is exited.
func (s *BaseSQLiteParserListener) ExitModule_name(ctx *Module_nameContext) {}

// EnterPragma_name is called when production pragma_name is entered.
func (s *BaseSQLiteParserListener) EnterPragma_name(ctx *Pragma_nameContext) {}

// ExitPragma_name is called when production pragma_name is exited.
func (s *BaseSQLiteParserListener) ExitPragma_name(ctx *Pragma_nameContext) {}

// EnterSavepoint_name is called when production savepoint_name is entered.
func (s *BaseSQLiteParserListener) EnterSavepoint_name(ctx *Savepoint_nameContext) {}

// ExitSavepoint_name is called when production savepoint_name is exited.
func (s *BaseSQLiteParserListener) ExitSavepoint_name(ctx *Savepoint_nameContext) {}

// EnterTable_alias is called when production table_alias is entered.
func (s *BaseSQLiteParserListener) EnterTable_alias(ctx *Table_aliasContext) {}

// ExitTable_alias is called when production table_alias is exited.
func (s *BaseSQLiteParserListener) ExitTable_alias(ctx *Table_aliasContext) {}

// EnterTransaction_name is called when production transaction_name is entered.
func (s *BaseSQLiteParserListener) EnterTransaction_name(ctx *Transaction_nameContext) {}

// ExitTransaction_name is called when production transaction_name is exited.
func (s *BaseSQLiteParserListener) ExitTransaction_name(ctx *Transaction_nameContext) {}

// EnterWindow_name is called when production window_name is entered.
func (s *BaseSQLiteParserListener) EnterWindow_name(ctx *Window_nameContext) {}

// ExitWindow_name is called when production window_name is exited.
func (s *BaseSQLiteParserListener) ExitWindow_name(ctx *Window_nameContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseSQLiteParserListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseSQLiteParserListener) ExitAlias(ctx *AliasContext) {}

// EnterFilename is called when production filename is entered.
func (s *BaseSQLiteParserListener) EnterFilename(ctx *FilenameContext) {}

// ExitFilename is called when production filename is exited.
func (s *BaseSQLiteParserListener) ExitFilename(ctx *FilenameContext) {}

// EnterBase_window_name is called when production base_window_name is entered.
func (s *BaseSQLiteParserListener) EnterBase_window_name(ctx *Base_window_nameContext) {}

// ExitBase_window_name is called when production base_window_name is exited.
func (s *BaseSQLiteParserListener) ExitBase_window_name(ctx *Base_window_nameContext) {}

// EnterSimple_func is called when production simple_func is entered.
func (s *BaseSQLiteParserListener) EnterSimple_func(ctx *Simple_funcContext) {}

// ExitSimple_func is called when production simple_func is exited.
func (s *BaseSQLiteParserListener) ExitSimple_func(ctx *Simple_funcContext) {}

// EnterAggregate_func is called when production aggregate_func is entered.
func (s *BaseSQLiteParserListener) EnterAggregate_func(ctx *Aggregate_funcContext) {}

// ExitAggregate_func is called when production aggregate_func is exited.
func (s *BaseSQLiteParserListener) ExitAggregate_func(ctx *Aggregate_funcContext) {}

// EnterWindow_func is called when production window_func is entered.
func (s *BaseSQLiteParserListener) EnterWindow_func(ctx *Window_funcContext) {}

// ExitWindow_func is called when production window_func is exited.
func (s *BaseSQLiteParserListener) ExitWindow_func(ctx *Window_funcContext) {}

// EnterTable_function_name is called when production table_function_name is entered.
func (s *BaseSQLiteParserListener) EnterTable_function_name(ctx *Table_function_nameContext) {}

// ExitTable_function_name is called when production table_function_name is exited.
func (s *BaseSQLiteParserListener) ExitTable_function_name(ctx *Table_function_nameContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSQLiteParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSQLiteParserListener) ExitIdentifier(ctx *IdentifierContext) {}
