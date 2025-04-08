// Code generated from ./sqlite/SQLiteParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // SQLiteParser

import "github.com/antlr4-go/antlr/v4"

// SQLiteParserListener is a complete listener for a parse tree produced by SQLiteParser.
type SQLiteParserListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterSql_stmt_list is called when entering the sql_stmt_list production.
	EnterSql_stmt_list(c *Sql_stmt_listContext)

	// EnterSql_stmt is called when entering the sql_stmt production.
	EnterSql_stmt(c *Sql_stmtContext)

	// EnterAlter_table_stmt is called when entering the alter_table_stmt production.
	EnterAlter_table_stmt(c *Alter_table_stmtContext)

	// EnterAnalyze_stmt is called when entering the analyze_stmt production.
	EnterAnalyze_stmt(c *Analyze_stmtContext)

	// EnterAttach_stmt is called when entering the attach_stmt production.
	EnterAttach_stmt(c *Attach_stmtContext)

	// EnterBegin_stmt is called when entering the begin_stmt production.
	EnterBegin_stmt(c *Begin_stmtContext)

	// EnterCommit_stmt is called when entering the commit_stmt production.
	EnterCommit_stmt(c *Commit_stmtContext)

	// EnterRollback_stmt is called when entering the rollback_stmt production.
	EnterRollback_stmt(c *Rollback_stmtContext)

	// EnterSavepoint_stmt is called when entering the savepoint_stmt production.
	EnterSavepoint_stmt(c *Savepoint_stmtContext)

	// EnterRelease_stmt is called when entering the release_stmt production.
	EnterRelease_stmt(c *Release_stmtContext)

	// EnterCreate_index_stmt is called when entering the create_index_stmt production.
	EnterCreate_index_stmt(c *Create_index_stmtContext)

	// EnterIndexed_column is called when entering the indexed_column production.
	EnterIndexed_column(c *Indexed_columnContext)

	// EnterCreate_table_stmt is called when entering the create_table_stmt production.
	EnterCreate_table_stmt(c *Create_table_stmtContext)

	// EnterColumn_def is called when entering the column_def production.
	EnterColumn_def(c *Column_defContext)

	// EnterType_name is called when entering the type_name production.
	EnterType_name(c *Type_nameContext)

	// EnterColumn_constraint is called when entering the column_constraint production.
	EnterColumn_constraint(c *Column_constraintContext)

	// EnterSigned_number is called when entering the signed_number production.
	EnterSigned_number(c *Signed_numberContext)

	// EnterTable_constraint is called when entering the table_constraint production.
	EnterTable_constraint(c *Table_constraintContext)

	// EnterForeign_key_clause is called when entering the foreign_key_clause production.
	EnterForeign_key_clause(c *Foreign_key_clauseContext)

	// EnterConflict_clause is called when entering the conflict_clause production.
	EnterConflict_clause(c *Conflict_clauseContext)

	// EnterCreate_trigger_stmt is called when entering the create_trigger_stmt production.
	EnterCreate_trigger_stmt(c *Create_trigger_stmtContext)

	// EnterCreate_view_stmt is called when entering the create_view_stmt production.
	EnterCreate_view_stmt(c *Create_view_stmtContext)

	// EnterCreate_virtual_table_stmt is called when entering the create_virtual_table_stmt production.
	EnterCreate_virtual_table_stmt(c *Create_virtual_table_stmtContext)

	// EnterDelete_stmt is called when entering the delete_stmt production.
	EnterDelete_stmt(c *Delete_stmtContext)

	// EnterDetach_stmt is called when entering the detach_stmt production.
	EnterDetach_stmt(c *Detach_stmtContext)

	// EnterDrop_stmt is called when entering the drop_stmt production.
	EnterDrop_stmt(c *Drop_stmtContext)

	// EnterExpr_case is called when entering the expr_case production.
	EnterExpr_case(c *Expr_caseContext)

	// EnterExpr_arithmetic is called when entering the expr_arithmetic production.
	EnterExpr_arithmetic(c *Expr_arithmeticContext)

	// EnterExpr_aggregate_func is called when entering the expr_aggregate_func production.
	EnterExpr_aggregate_func(c *Expr_aggregate_funcContext)

	// EnterExpr_json_extract_string is called when entering the expr_json_extract_string production.
	EnterExpr_json_extract_string(c *Expr_json_extract_stringContext)

	// EnterExpr_raise is called when entering the expr_raise production.
	EnterExpr_raise(c *Expr_raiseContext)

	// EnterExpr_bool is called when entering the expr_bool production.
	EnterExpr_bool(c *Expr_boolContext)

	// EnterExpr_is is called when entering the expr_is production.
	EnterExpr_is(c *Expr_isContext)

	// EnterExpr_concat is called when entering the expr_concat production.
	EnterExpr_concat(c *Expr_concatContext)

	// EnterExpr_list is called when entering the expr_list production.
	EnterExpr_list(c *Expr_listContext)

	// EnterExpr_in is called when entering the expr_in production.
	EnterExpr_in(c *Expr_inContext)

	// EnterExpr_collate is called when entering the expr_collate production.
	EnterExpr_collate(c *Expr_collateContext)

	// EnterExpr_modulo is called when entering the expr_modulo production.
	EnterExpr_modulo(c *Expr_moduloContext)

	// EnterExpr_qualified_column_name is called when entering the expr_qualified_column_name production.
	EnterExpr_qualified_column_name(c *Expr_qualified_column_nameContext)

	// EnterExpr_match is called when entering the expr_match production.
	EnterExpr_match(c *Expr_matchContext)

	// EnterExpr_like is called when entering the expr_like production.
	EnterExpr_like(c *Expr_likeContext)

	// EnterExpr_null_comp is called when entering the expr_null_comp production.
	EnterExpr_null_comp(c *Expr_null_compContext)

	// EnterExpr_json_extract_json is called when entering the expr_json_extract_json production.
	EnterExpr_json_extract_json(c *Expr_json_extract_jsonContext)

	// EnterExpr_window_func is called when entering the expr_window_func production.
	EnterExpr_window_func(c *Expr_window_funcContext)

	// EnterExpr_exists_select is called when entering the expr_exists_select production.
	EnterExpr_exists_select(c *Expr_exists_selectContext)

	// EnterExpr_comparison is called when entering the expr_comparison production.
	EnterExpr_comparison(c *Expr_comparisonContext)

	// EnterExpr_literal is called when entering the expr_literal production.
	EnterExpr_literal(c *Expr_literalContext)

	// EnterExpr_cast is called when entering the expr_cast production.
	EnterExpr_cast(c *Expr_castContext)

	// EnterExpr_string_op is called when entering the expr_string_op production.
	EnterExpr_string_op(c *Expr_string_opContext)

	// EnterExpr_between is called when entering the expr_between production.
	EnterExpr_between(c *Expr_betweenContext)

	// EnterExpr_bitwise is called when entering the expr_bitwise production.
	EnterExpr_bitwise(c *Expr_bitwiseContext)

	// EnterExpr_simple_func is called when entering the expr_simple_func production.
	EnterExpr_simple_func(c *Expr_simple_funcContext)

	// EnterExpr_unary is called when entering the expr_unary production.
	EnterExpr_unary(c *Expr_unaryContext)

	// EnterExpr_bind is called when entering the expr_bind production.
	EnterExpr_bind(c *Expr_bindContext)

	// EnterRaise_function is called when entering the raise_function production.
	EnterRaise_function(c *Raise_functionContext)

	// EnterLiteral_value is called when entering the literal_value production.
	EnterLiteral_value(c *Literal_valueContext)

	// EnterValue_row is called when entering the value_row production.
	EnterValue_row(c *Value_rowContext)

	// EnterValues_clause is called when entering the values_clause production.
	EnterValues_clause(c *Values_clauseContext)

	// EnterInsert_stmt is called when entering the insert_stmt production.
	EnterInsert_stmt(c *Insert_stmtContext)

	// EnterReturning_clause is called when entering the returning_clause production.
	EnterReturning_clause(c *Returning_clauseContext)

	// EnterUpsert_clause is called when entering the upsert_clause production.
	EnterUpsert_clause(c *Upsert_clauseContext)

	// EnterPragma_stmt is called when entering the pragma_stmt production.
	EnterPragma_stmt(c *Pragma_stmtContext)

	// EnterPragma_value is called when entering the pragma_value production.
	EnterPragma_value(c *Pragma_valueContext)

	// EnterReindex_stmt is called when entering the reindex_stmt production.
	EnterReindex_stmt(c *Reindex_stmtContext)

	// EnterSelect_stmt is called when entering the select_stmt production.
	EnterSelect_stmt(c *Select_stmtContext)

	// EnterSelect_core is called when entering the select_core production.
	EnterSelect_core(c *Select_coreContext)

	// EnterResult_column is called when entering the result_column production.
	EnterResult_column(c *Result_columnContext)

	// EnterFrom_item is called when entering the from_item production.
	EnterFrom_item(c *From_itemContext)

	// EnterJoin_operator is called when entering the join_operator production.
	EnterJoin_operator(c *Join_operatorContext)

	// EnterJoin_constraint is called when entering the join_constraint production.
	EnterJoin_constraint(c *Join_constraintContext)

	// EnterTable_or_subquery is called when entering the table_or_subquery production.
	EnterTable_or_subquery(c *Table_or_subqueryContext)

	// EnterCompound_select is called when entering the compound_select production.
	EnterCompound_select(c *Compound_selectContext)

	// EnterCompound_operator is called when entering the compound_operator production.
	EnterCompound_operator(c *Compound_operatorContext)

	// EnterUpdate_stmt is called when entering the update_stmt production.
	EnterUpdate_stmt(c *Update_stmtContext)

	// EnterColumn_name_list is called when entering the column_name_list production.
	EnterColumn_name_list(c *Column_name_listContext)

	// EnterQualified_table_name is called when entering the qualified_table_name production.
	EnterQualified_table_name(c *Qualified_table_nameContext)

	// EnterVacuum_stmt is called when entering the vacuum_stmt production.
	EnterVacuum_stmt(c *Vacuum_stmtContext)

	// EnterFilter_clause is called when entering the filter_clause production.
	EnterFilter_clause(c *Filter_clauseContext)

	// EnterWindow_defn is called when entering the window_defn production.
	EnterWindow_defn(c *Window_defnContext)

	// EnterOver_clause is called when entering the over_clause production.
	EnterOver_clause(c *Over_clauseContext)

	// EnterFrame_spec is called when entering the frame_spec production.
	EnterFrame_spec(c *Frame_specContext)

	// EnterFrame_clause is called when entering the frame_clause production.
	EnterFrame_clause(c *Frame_clauseContext)

	// EnterWith_clause is called when entering the with_clause production.
	EnterWith_clause(c *With_clauseContext)

	// EnterCommon_table_expression is called when entering the common_table_expression production.
	EnterCommon_table_expression(c *Common_table_expressionContext)

	// EnterWhere_stmt is called when entering the where_stmt production.
	EnterWhere_stmt(c *Where_stmtContext)

	// EnterOrder_by_stmt is called when entering the order_by_stmt production.
	EnterOrder_by_stmt(c *Order_by_stmtContext)

	// EnterGroup_by_stmt is called when entering the group_by_stmt production.
	EnterGroup_by_stmt(c *Group_by_stmtContext)

	// EnterWindow_stmt is called when entering the window_stmt production.
	EnterWindow_stmt(c *Window_stmtContext)

	// EnterLimit_stmt is called when entering the limit_stmt production.
	EnterLimit_stmt(c *Limit_stmtContext)

	// EnterOrdering_term is called when entering the ordering_term production.
	EnterOrdering_term(c *Ordering_termContext)

	// EnterAsc_desc is called when entering the asc_desc production.
	EnterAsc_desc(c *Asc_descContext)

	// EnterFrame_left is called when entering the frame_left production.
	EnterFrame_left(c *Frame_leftContext)

	// EnterFrame_right is called when entering the frame_right production.
	EnterFrame_right(c *Frame_rightContext)

	// EnterFrame_single is called when entering the frame_single production.
	EnterFrame_single(c *Frame_singleContext)

	// EnterOffset is called when entering the offset production.
	EnterOffset(c *OffsetContext)

	// EnterDefault_value is called when entering the default_value production.
	EnterDefault_value(c *Default_valueContext)

	// EnterPartition_by is called when entering the partition_by production.
	EnterPartition_by(c *Partition_byContext)

	// EnterOrder_by_expr is called when entering the order_by_expr production.
	EnterOrder_by_expr(c *Order_by_exprContext)

	// EnterOrder_by_expr_asc_desc is called when entering the order_by_expr_asc_desc production.
	EnterOrder_by_expr_asc_desc(c *Order_by_expr_asc_descContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterError_message is called when entering the error_message production.
	EnterError_message(c *Error_messageContext)

	// EnterModule_argument is called when entering the module_argument production.
	EnterModule_argument(c *Module_argumentContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterFunction_name is called when entering the function_name production.
	EnterFunction_name(c *Function_nameContext)

	// EnterSchema_name is called when entering the schema_name production.
	EnterSchema_name(c *Schema_nameContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterTable_or_index_name is called when entering the table_or_index_name production.
	EnterTable_or_index_name(c *Table_or_index_nameContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterCollation_name is called when entering the collation_name production.
	EnterCollation_name(c *Collation_nameContext)

	// EnterForeign_table is called when entering the foreign_table production.
	EnterForeign_table(c *Foreign_tableContext)

	// EnterIndex_name is called when entering the index_name production.
	EnterIndex_name(c *Index_nameContext)

	// EnterTrigger_name is called when entering the trigger_name production.
	EnterTrigger_name(c *Trigger_nameContext)

	// EnterView_name is called when entering the view_name production.
	EnterView_name(c *View_nameContext)

	// EnterModule_name is called when entering the module_name production.
	EnterModule_name(c *Module_nameContext)

	// EnterPragma_name is called when entering the pragma_name production.
	EnterPragma_name(c *Pragma_nameContext)

	// EnterSavepoint_name is called when entering the savepoint_name production.
	EnterSavepoint_name(c *Savepoint_nameContext)

	// EnterTable_alias is called when entering the table_alias production.
	EnterTable_alias(c *Table_aliasContext)

	// EnterTransaction_name is called when entering the transaction_name production.
	EnterTransaction_name(c *Transaction_nameContext)

	// EnterWindow_name is called when entering the window_name production.
	EnterWindow_name(c *Window_nameContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterFilename is called when entering the filename production.
	EnterFilename(c *FilenameContext)

	// EnterBase_window_name is called when entering the base_window_name production.
	EnterBase_window_name(c *Base_window_nameContext)

	// EnterSimple_func is called when entering the simple_func production.
	EnterSimple_func(c *Simple_funcContext)

	// EnterAggregate_func is called when entering the aggregate_func production.
	EnterAggregate_func(c *Aggregate_funcContext)

	// EnterWindow_func is called when entering the window_func production.
	EnterWindow_func(c *Window_funcContext)

	// EnterTable_function_name is called when entering the table_function_name production.
	EnterTable_function_name(c *Table_function_nameContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitSql_stmt_list is called when exiting the sql_stmt_list production.
	ExitSql_stmt_list(c *Sql_stmt_listContext)

	// ExitSql_stmt is called when exiting the sql_stmt production.
	ExitSql_stmt(c *Sql_stmtContext)

	// ExitAlter_table_stmt is called when exiting the alter_table_stmt production.
	ExitAlter_table_stmt(c *Alter_table_stmtContext)

	// ExitAnalyze_stmt is called when exiting the analyze_stmt production.
	ExitAnalyze_stmt(c *Analyze_stmtContext)

	// ExitAttach_stmt is called when exiting the attach_stmt production.
	ExitAttach_stmt(c *Attach_stmtContext)

	// ExitBegin_stmt is called when exiting the begin_stmt production.
	ExitBegin_stmt(c *Begin_stmtContext)

	// ExitCommit_stmt is called when exiting the commit_stmt production.
	ExitCommit_stmt(c *Commit_stmtContext)

	// ExitRollback_stmt is called when exiting the rollback_stmt production.
	ExitRollback_stmt(c *Rollback_stmtContext)

	// ExitSavepoint_stmt is called when exiting the savepoint_stmt production.
	ExitSavepoint_stmt(c *Savepoint_stmtContext)

	// ExitRelease_stmt is called when exiting the release_stmt production.
	ExitRelease_stmt(c *Release_stmtContext)

	// ExitCreate_index_stmt is called when exiting the create_index_stmt production.
	ExitCreate_index_stmt(c *Create_index_stmtContext)

	// ExitIndexed_column is called when exiting the indexed_column production.
	ExitIndexed_column(c *Indexed_columnContext)

	// ExitCreate_table_stmt is called when exiting the create_table_stmt production.
	ExitCreate_table_stmt(c *Create_table_stmtContext)

	// ExitColumn_def is called when exiting the column_def production.
	ExitColumn_def(c *Column_defContext)

	// ExitType_name is called when exiting the type_name production.
	ExitType_name(c *Type_nameContext)

	// ExitColumn_constraint is called when exiting the column_constraint production.
	ExitColumn_constraint(c *Column_constraintContext)

	// ExitSigned_number is called when exiting the signed_number production.
	ExitSigned_number(c *Signed_numberContext)

	// ExitTable_constraint is called when exiting the table_constraint production.
	ExitTable_constraint(c *Table_constraintContext)

	// ExitForeign_key_clause is called when exiting the foreign_key_clause production.
	ExitForeign_key_clause(c *Foreign_key_clauseContext)

	// ExitConflict_clause is called when exiting the conflict_clause production.
	ExitConflict_clause(c *Conflict_clauseContext)

	// ExitCreate_trigger_stmt is called when exiting the create_trigger_stmt production.
	ExitCreate_trigger_stmt(c *Create_trigger_stmtContext)

	// ExitCreate_view_stmt is called when exiting the create_view_stmt production.
	ExitCreate_view_stmt(c *Create_view_stmtContext)

	// ExitCreate_virtual_table_stmt is called when exiting the create_virtual_table_stmt production.
	ExitCreate_virtual_table_stmt(c *Create_virtual_table_stmtContext)

	// ExitDelete_stmt is called when exiting the delete_stmt production.
	ExitDelete_stmt(c *Delete_stmtContext)

	// ExitDetach_stmt is called when exiting the detach_stmt production.
	ExitDetach_stmt(c *Detach_stmtContext)

	// ExitDrop_stmt is called when exiting the drop_stmt production.
	ExitDrop_stmt(c *Drop_stmtContext)

	// ExitExpr_case is called when exiting the expr_case production.
	ExitExpr_case(c *Expr_caseContext)

	// ExitExpr_arithmetic is called when exiting the expr_arithmetic production.
	ExitExpr_arithmetic(c *Expr_arithmeticContext)

	// ExitExpr_aggregate_func is called when exiting the expr_aggregate_func production.
	ExitExpr_aggregate_func(c *Expr_aggregate_funcContext)

	// ExitExpr_json_extract_string is called when exiting the expr_json_extract_string production.
	ExitExpr_json_extract_string(c *Expr_json_extract_stringContext)

	// ExitExpr_raise is called when exiting the expr_raise production.
	ExitExpr_raise(c *Expr_raiseContext)

	// ExitExpr_bool is called when exiting the expr_bool production.
	ExitExpr_bool(c *Expr_boolContext)

	// ExitExpr_is is called when exiting the expr_is production.
	ExitExpr_is(c *Expr_isContext)

	// ExitExpr_concat is called when exiting the expr_concat production.
	ExitExpr_concat(c *Expr_concatContext)

	// ExitExpr_list is called when exiting the expr_list production.
	ExitExpr_list(c *Expr_listContext)

	// ExitExpr_in is called when exiting the expr_in production.
	ExitExpr_in(c *Expr_inContext)

	// ExitExpr_collate is called when exiting the expr_collate production.
	ExitExpr_collate(c *Expr_collateContext)

	// ExitExpr_modulo is called when exiting the expr_modulo production.
	ExitExpr_modulo(c *Expr_moduloContext)

	// ExitExpr_qualified_column_name is called when exiting the expr_qualified_column_name production.
	ExitExpr_qualified_column_name(c *Expr_qualified_column_nameContext)

	// ExitExpr_match is called when exiting the expr_match production.
	ExitExpr_match(c *Expr_matchContext)

	// ExitExpr_like is called when exiting the expr_like production.
	ExitExpr_like(c *Expr_likeContext)

	// ExitExpr_null_comp is called when exiting the expr_null_comp production.
	ExitExpr_null_comp(c *Expr_null_compContext)

	// ExitExpr_json_extract_json is called when exiting the expr_json_extract_json production.
	ExitExpr_json_extract_json(c *Expr_json_extract_jsonContext)

	// ExitExpr_window_func is called when exiting the expr_window_func production.
	ExitExpr_window_func(c *Expr_window_funcContext)

	// ExitExpr_exists_select is called when exiting the expr_exists_select production.
	ExitExpr_exists_select(c *Expr_exists_selectContext)

	// ExitExpr_comparison is called when exiting the expr_comparison production.
	ExitExpr_comparison(c *Expr_comparisonContext)

	// ExitExpr_literal is called when exiting the expr_literal production.
	ExitExpr_literal(c *Expr_literalContext)

	// ExitExpr_cast is called when exiting the expr_cast production.
	ExitExpr_cast(c *Expr_castContext)

	// ExitExpr_string_op is called when exiting the expr_string_op production.
	ExitExpr_string_op(c *Expr_string_opContext)

	// ExitExpr_between is called when exiting the expr_between production.
	ExitExpr_between(c *Expr_betweenContext)

	// ExitExpr_bitwise is called when exiting the expr_bitwise production.
	ExitExpr_bitwise(c *Expr_bitwiseContext)

	// ExitExpr_simple_func is called when exiting the expr_simple_func production.
	ExitExpr_simple_func(c *Expr_simple_funcContext)

	// ExitExpr_unary is called when exiting the expr_unary production.
	ExitExpr_unary(c *Expr_unaryContext)

	// ExitExpr_bind is called when exiting the expr_bind production.
	ExitExpr_bind(c *Expr_bindContext)

	// ExitRaise_function is called when exiting the raise_function production.
	ExitRaise_function(c *Raise_functionContext)

	// ExitLiteral_value is called when exiting the literal_value production.
	ExitLiteral_value(c *Literal_valueContext)

	// ExitValue_row is called when exiting the value_row production.
	ExitValue_row(c *Value_rowContext)

	// ExitValues_clause is called when exiting the values_clause production.
	ExitValues_clause(c *Values_clauseContext)

	// ExitInsert_stmt is called when exiting the insert_stmt production.
	ExitInsert_stmt(c *Insert_stmtContext)

	// ExitReturning_clause is called when exiting the returning_clause production.
	ExitReturning_clause(c *Returning_clauseContext)

	// ExitUpsert_clause is called when exiting the upsert_clause production.
	ExitUpsert_clause(c *Upsert_clauseContext)

	// ExitPragma_stmt is called when exiting the pragma_stmt production.
	ExitPragma_stmt(c *Pragma_stmtContext)

	// ExitPragma_value is called when exiting the pragma_value production.
	ExitPragma_value(c *Pragma_valueContext)

	// ExitReindex_stmt is called when exiting the reindex_stmt production.
	ExitReindex_stmt(c *Reindex_stmtContext)

	// ExitSelect_stmt is called when exiting the select_stmt production.
	ExitSelect_stmt(c *Select_stmtContext)

	// ExitSelect_core is called when exiting the select_core production.
	ExitSelect_core(c *Select_coreContext)

	// ExitResult_column is called when exiting the result_column production.
	ExitResult_column(c *Result_columnContext)

	// ExitFrom_item is called when exiting the from_item production.
	ExitFrom_item(c *From_itemContext)

	// ExitJoin_operator is called when exiting the join_operator production.
	ExitJoin_operator(c *Join_operatorContext)

	// ExitJoin_constraint is called when exiting the join_constraint production.
	ExitJoin_constraint(c *Join_constraintContext)

	// ExitTable_or_subquery is called when exiting the table_or_subquery production.
	ExitTable_or_subquery(c *Table_or_subqueryContext)

	// ExitCompound_select is called when exiting the compound_select production.
	ExitCompound_select(c *Compound_selectContext)

	// ExitCompound_operator is called when exiting the compound_operator production.
	ExitCompound_operator(c *Compound_operatorContext)

	// ExitUpdate_stmt is called when exiting the update_stmt production.
	ExitUpdate_stmt(c *Update_stmtContext)

	// ExitColumn_name_list is called when exiting the column_name_list production.
	ExitColumn_name_list(c *Column_name_listContext)

	// ExitQualified_table_name is called when exiting the qualified_table_name production.
	ExitQualified_table_name(c *Qualified_table_nameContext)

	// ExitVacuum_stmt is called when exiting the vacuum_stmt production.
	ExitVacuum_stmt(c *Vacuum_stmtContext)

	// ExitFilter_clause is called when exiting the filter_clause production.
	ExitFilter_clause(c *Filter_clauseContext)

	// ExitWindow_defn is called when exiting the window_defn production.
	ExitWindow_defn(c *Window_defnContext)

	// ExitOver_clause is called when exiting the over_clause production.
	ExitOver_clause(c *Over_clauseContext)

	// ExitFrame_spec is called when exiting the frame_spec production.
	ExitFrame_spec(c *Frame_specContext)

	// ExitFrame_clause is called when exiting the frame_clause production.
	ExitFrame_clause(c *Frame_clauseContext)

	// ExitWith_clause is called when exiting the with_clause production.
	ExitWith_clause(c *With_clauseContext)

	// ExitCommon_table_expression is called when exiting the common_table_expression production.
	ExitCommon_table_expression(c *Common_table_expressionContext)

	// ExitWhere_stmt is called when exiting the where_stmt production.
	ExitWhere_stmt(c *Where_stmtContext)

	// ExitOrder_by_stmt is called when exiting the order_by_stmt production.
	ExitOrder_by_stmt(c *Order_by_stmtContext)

	// ExitGroup_by_stmt is called when exiting the group_by_stmt production.
	ExitGroup_by_stmt(c *Group_by_stmtContext)

	// ExitWindow_stmt is called when exiting the window_stmt production.
	ExitWindow_stmt(c *Window_stmtContext)

	// ExitLimit_stmt is called when exiting the limit_stmt production.
	ExitLimit_stmt(c *Limit_stmtContext)

	// ExitOrdering_term is called when exiting the ordering_term production.
	ExitOrdering_term(c *Ordering_termContext)

	// ExitAsc_desc is called when exiting the asc_desc production.
	ExitAsc_desc(c *Asc_descContext)

	// ExitFrame_left is called when exiting the frame_left production.
	ExitFrame_left(c *Frame_leftContext)

	// ExitFrame_right is called when exiting the frame_right production.
	ExitFrame_right(c *Frame_rightContext)

	// ExitFrame_single is called when exiting the frame_single production.
	ExitFrame_single(c *Frame_singleContext)

	// ExitOffset is called when exiting the offset production.
	ExitOffset(c *OffsetContext)

	// ExitDefault_value is called when exiting the default_value production.
	ExitDefault_value(c *Default_valueContext)

	// ExitPartition_by is called when exiting the partition_by production.
	ExitPartition_by(c *Partition_byContext)

	// ExitOrder_by_expr is called when exiting the order_by_expr production.
	ExitOrder_by_expr(c *Order_by_exprContext)

	// ExitOrder_by_expr_asc_desc is called when exiting the order_by_expr_asc_desc production.
	ExitOrder_by_expr_asc_desc(c *Order_by_expr_asc_descContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitError_message is called when exiting the error_message production.
	ExitError_message(c *Error_messageContext)

	// ExitModule_argument is called when exiting the module_argument production.
	ExitModule_argument(c *Module_argumentContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitFunction_name is called when exiting the function_name production.
	ExitFunction_name(c *Function_nameContext)

	// ExitSchema_name is called when exiting the schema_name production.
	ExitSchema_name(c *Schema_nameContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitTable_or_index_name is called when exiting the table_or_index_name production.
	ExitTable_or_index_name(c *Table_or_index_nameContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitCollation_name is called when exiting the collation_name production.
	ExitCollation_name(c *Collation_nameContext)

	// ExitForeign_table is called when exiting the foreign_table production.
	ExitForeign_table(c *Foreign_tableContext)

	// ExitIndex_name is called when exiting the index_name production.
	ExitIndex_name(c *Index_nameContext)

	// ExitTrigger_name is called when exiting the trigger_name production.
	ExitTrigger_name(c *Trigger_nameContext)

	// ExitView_name is called when exiting the view_name production.
	ExitView_name(c *View_nameContext)

	// ExitModule_name is called when exiting the module_name production.
	ExitModule_name(c *Module_nameContext)

	// ExitPragma_name is called when exiting the pragma_name production.
	ExitPragma_name(c *Pragma_nameContext)

	// ExitSavepoint_name is called when exiting the savepoint_name production.
	ExitSavepoint_name(c *Savepoint_nameContext)

	// ExitTable_alias is called when exiting the table_alias production.
	ExitTable_alias(c *Table_aliasContext)

	// ExitTransaction_name is called when exiting the transaction_name production.
	ExitTransaction_name(c *Transaction_nameContext)

	// ExitWindow_name is called when exiting the window_name production.
	ExitWindow_name(c *Window_nameContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitFilename is called when exiting the filename production.
	ExitFilename(c *FilenameContext)

	// ExitBase_window_name is called when exiting the base_window_name production.
	ExitBase_window_name(c *Base_window_nameContext)

	// ExitSimple_func is called when exiting the simple_func production.
	ExitSimple_func(c *Simple_funcContext)

	// ExitAggregate_func is called when exiting the aggregate_func production.
	ExitAggregate_func(c *Aggregate_funcContext)

	// ExitWindow_func is called when exiting the window_func production.
	ExitWindow_func(c *Window_funcContext)

	// ExitTable_function_name is called when exiting the table_function_name production.
	ExitTable_function_name(c *Table_function_nameContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)
}
