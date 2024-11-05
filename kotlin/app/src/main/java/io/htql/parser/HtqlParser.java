// Generated from Htql.g4 by ANTLR 4.13.2
package io.htql.parser;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class HtqlParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, SELECT=2, FROM=3, WHERE=4, AND=5, OR=6, IS=7, NOT=8, NULL=9, STAR=10, 
		COMMA=11, DOT=12, EQUALS=13, IDENTIFIER=14, STRING=15, WS=16, URL=17;
	public static final int
		RULE_query = 0, RULE_selectStmt = 1, RULE_elementList = 2, RULE_fromStmt = 3, 
		RULE_whereStmt = 4, RULE_conditionExpr = 5, RULE_condition = 6, RULE_attributeExpr = 7, 
		RULE_attributeNullCheck = 8, RULE_logicalOp = 9;
	private static String[] makeRuleNames() {
		return new String[] {
			"query", "selectStmt", "elementList", "fromStmt", "whereStmt", "conditionExpr", 
			"condition", "attributeExpr", "attributeNullCheck", "logicalOp"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'attributes'", "'SELECT'", "'FROM'", "'WHERE'", "'AND'", "'OR'", 
			"'IS'", "'NOT'", "'NULL'", "'*'", "','", "'.'", "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, "SELECT", "FROM", "WHERE", "AND", "OR", "IS", "NOT", "NULL", 
			"STAR", "COMMA", "DOT", "EQUALS", "IDENTIFIER", "STRING", "WS", "URL"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Htql.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public HtqlParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class QueryContext extends ParserRuleContext {
		public SelectStmtContext selectStmt() {
			return getRuleContext(SelectStmtContext.class,0);
		}
		public FromStmtContext fromStmt() {
			return getRuleContext(FromStmtContext.class,0);
		}
		public WhereStmtContext whereStmt() {
			return getRuleContext(WhereStmtContext.class,0);
		}
		public QueryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_query; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).enterQuery(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).exitQuery(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HtqlVisitor ) return ((HtqlVisitor<? extends T>)visitor).visitQuery(this);
			else return visitor.visitChildren(this);
		}
	}

	public final QueryContext query() throws RecognitionException {
		QueryContext _localctx = new QueryContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_query);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(20);
			selectStmt();
			setState(21);
			fromStmt();
			setState(23);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHERE) {
				{
				setState(22);
				whereStmt();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SelectStmtContext extends ParserRuleContext {
		public TerminalNode SELECT() { return getToken(HtqlParser.SELECT, 0); }
		public TerminalNode STAR() { return getToken(HtqlParser.STAR, 0); }
		public ElementListContext elementList() {
			return getRuleContext(ElementListContext.class,0);
		}
		public SelectStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectStmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).enterSelectStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).exitSelectStmt(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HtqlVisitor ) return ((HtqlVisitor<? extends T>)visitor).visitSelectStmt(this);
			else return visitor.visitChildren(this);
		}
	}

	public final SelectStmtContext selectStmt() throws RecognitionException {
		SelectStmtContext _localctx = new SelectStmtContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_selectStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(25);
			match(SELECT);
			setState(28);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STAR:
				{
				setState(26);
				match(STAR);
				}
				break;
			case IDENTIFIER:
				{
				setState(27);
				elementList();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ElementListContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(HtqlParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(HtqlParser.IDENTIFIER, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(HtqlParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(HtqlParser.COMMA, i);
		}
		public ElementListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elementList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).enterElementList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).exitElementList(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HtqlVisitor ) return ((HtqlVisitor<? extends T>)visitor).visitElementList(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ElementListContext elementList() throws RecognitionException {
		ElementListContext _localctx = new ElementListContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_elementList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(30);
			match(IDENTIFIER);
			setState(35);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(31);
				match(COMMA);
				setState(32);
				match(IDENTIFIER);
				}
				}
				setState(37);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FromStmtContext extends ParserRuleContext {
		public TerminalNode FROM() { return getToken(HtqlParser.FROM, 0); }
		public TerminalNode IDENTIFIER() { return getToken(HtqlParser.IDENTIFIER, 0); }
		public TerminalNode URL() { return getToken(HtqlParser.URL, 0); }
		public FromStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fromStmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).enterFromStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).exitFromStmt(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HtqlVisitor ) return ((HtqlVisitor<? extends T>)visitor).visitFromStmt(this);
			else return visitor.visitChildren(this);
		}
	}

	public final FromStmtContext fromStmt() throws RecognitionException {
		FromStmtContext _localctx = new FromStmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_fromStmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(38);
			match(FROM);
			setState(39);
			_la = _input.LA(1);
			if ( !(_la==IDENTIFIER || _la==URL) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class WhereStmtContext extends ParserRuleContext {
		public TerminalNode WHERE() { return getToken(HtqlParser.WHERE, 0); }
		public ConditionExprContext conditionExpr() {
			return getRuleContext(ConditionExprContext.class,0);
		}
		public WhereStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whereStmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).enterWhereStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).exitWhereStmt(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HtqlVisitor ) return ((HtqlVisitor<? extends T>)visitor).visitWhereStmt(this);
			else return visitor.visitChildren(this);
		}
	}

	public final WhereStmtContext whereStmt() throws RecognitionException {
		WhereStmtContext _localctx = new WhereStmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_whereStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(41);
			match(WHERE);
			setState(42);
			conditionExpr();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionExprContext extends ParserRuleContext {
		public List<ConditionContext> condition() {
			return getRuleContexts(ConditionContext.class);
		}
		public ConditionContext condition(int i) {
			return getRuleContext(ConditionContext.class,i);
		}
		public List<LogicalOpContext> logicalOp() {
			return getRuleContexts(LogicalOpContext.class);
		}
		public LogicalOpContext logicalOp(int i) {
			return getRuleContext(LogicalOpContext.class,i);
		}
		public ConditionExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditionExpr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).enterConditionExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).exitConditionExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HtqlVisitor ) return ((HtqlVisitor<? extends T>)visitor).visitConditionExpr(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ConditionExprContext conditionExpr() throws RecognitionException {
		ConditionExprContext _localctx = new ConditionExprContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_conditionExpr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			condition();
			setState(50);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND || _la==OR) {
				{
				{
				setState(45);
				logicalOp();
				setState(46);
				condition();
				}
				}
				setState(52);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConditionContext extends ParserRuleContext {
		public AttributeExprContext attributeExpr() {
			return getRuleContext(AttributeExprContext.class,0);
		}
		public AttributeNullCheckContext attributeNullCheck() {
			return getRuleContext(AttributeNullCheckContext.class,0);
		}
		public TerminalNode NOT() { return getToken(HtqlParser.NOT, 0); }
		public ConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condition; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).enterCondition(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).exitCondition(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HtqlVisitor ) return ((HtqlVisitor<? extends T>)visitor).visitCondition(this);
			else return visitor.visitChildren(this);
		}
	}

	public final ConditionContext condition() throws RecognitionException {
		ConditionContext _localctx = new ConditionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_condition);
		try {
			setState(57);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(53);
				attributeExpr();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(54);
				attributeNullCheck();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(55);
				match(NOT);
				setState(56);
				attributeExpr();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AttributeExprContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(HtqlParser.DOT, 0); }
		public TerminalNode IDENTIFIER() { return getToken(HtqlParser.IDENTIFIER, 0); }
		public TerminalNode EQUALS() { return getToken(HtqlParser.EQUALS, 0); }
		public TerminalNode STRING() { return getToken(HtqlParser.STRING, 0); }
		public AttributeExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributeExpr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).enterAttributeExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).exitAttributeExpr(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HtqlVisitor ) return ((HtqlVisitor<? extends T>)visitor).visitAttributeExpr(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributeExprContext attributeExpr() throws RecognitionException {
		AttributeExprContext _localctx = new AttributeExprContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_attributeExpr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(59);
			match(T__0);
			setState(60);
			match(DOT);
			setState(61);
			match(IDENTIFIER);
			setState(62);
			match(EQUALS);
			setState(63);
			match(STRING);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AttributeNullCheckContext extends ParserRuleContext {
		public TerminalNode IS() { return getToken(HtqlParser.IS, 0); }
		public TerminalNode NULL() { return getToken(HtqlParser.NULL, 0); }
		public TerminalNode NOT() { return getToken(HtqlParser.NOT, 0); }
		public AttributeNullCheckContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributeNullCheck; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).enterAttributeNullCheck(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).exitAttributeNullCheck(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HtqlVisitor ) return ((HtqlVisitor<? extends T>)visitor).visitAttributeNullCheck(this);
			else return visitor.visitChildren(this);
		}
	}

	public final AttributeNullCheckContext attributeNullCheck() throws RecognitionException {
		AttributeNullCheckContext _localctx = new AttributeNullCheckContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_attributeNullCheck);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(65);
			match(T__0);
			setState(66);
			match(IS);
			setState(68);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NOT) {
				{
				setState(67);
				match(NOT);
				}
			}

			setState(70);
			match(NULL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LogicalOpContext extends ParserRuleContext {
		public TerminalNode AND() { return getToken(HtqlParser.AND, 0); }
		public TerminalNode OR() { return getToken(HtqlParser.OR, 0); }
		public LogicalOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_logicalOp; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).enterLogicalOp(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof HtqlListener ) ((HtqlListener)listener).exitLogicalOp(this);
		}
		@Override
		public <T> T accept(ParseTreeVisitor<? extends T> visitor) {
			if ( visitor instanceof HtqlVisitor ) return ((HtqlVisitor<? extends T>)visitor).visitLogicalOp(this);
			else return visitor.visitChildren(this);
		}
	}

	public final LogicalOpContext logicalOp() throws RecognitionException {
		LogicalOpContext _localctx = new LogicalOpContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_logicalOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(72);
			_la = _input.LA(1);
			if ( !(_la==AND || _la==OR) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u0011K\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0001\u0000\u0001\u0000\u0001\u0000\u0003\u0000"+
		"\u0018\b\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001\u001d\b"+
		"\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0005\u0002\"\b\u0002\n\u0002"+
		"\f\u0002%\t\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005"+
		"\u00051\b\u0005\n\u0005\f\u00054\t\u0005\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0003\u0006:\b\u0006\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0003"+
		"\bE\b\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0000\u0000\n\u0000\u0002"+
		"\u0004\u0006\b\n\f\u000e\u0010\u0012\u0000\u0002\u0002\u0000\u000e\u000e"+
		"\u0011\u0011\u0001\u0000\u0005\u0006G\u0000\u0014\u0001\u0000\u0000\u0000"+
		"\u0002\u0019\u0001\u0000\u0000\u0000\u0004\u001e\u0001\u0000\u0000\u0000"+
		"\u0006&\u0001\u0000\u0000\u0000\b)\u0001\u0000\u0000\u0000\n,\u0001\u0000"+
		"\u0000\u0000\f9\u0001\u0000\u0000\u0000\u000e;\u0001\u0000\u0000\u0000"+
		"\u0010A\u0001\u0000\u0000\u0000\u0012H\u0001\u0000\u0000\u0000\u0014\u0015"+
		"\u0003\u0002\u0001\u0000\u0015\u0017\u0003\u0006\u0003\u0000\u0016\u0018"+
		"\u0003\b\u0004\u0000\u0017\u0016\u0001\u0000\u0000\u0000\u0017\u0018\u0001"+
		"\u0000\u0000\u0000\u0018\u0001\u0001\u0000\u0000\u0000\u0019\u001c\u0005"+
		"\u0002\u0000\u0000\u001a\u001d\u0005\n\u0000\u0000\u001b\u001d\u0003\u0004"+
		"\u0002\u0000\u001c\u001a\u0001\u0000\u0000\u0000\u001c\u001b\u0001\u0000"+
		"\u0000\u0000\u001d\u0003\u0001\u0000\u0000\u0000\u001e#\u0005\u000e\u0000"+
		"\u0000\u001f \u0005\u000b\u0000\u0000 \"\u0005\u000e\u0000\u0000!\u001f"+
		"\u0001\u0000\u0000\u0000\"%\u0001\u0000\u0000\u0000#!\u0001\u0000\u0000"+
		"\u0000#$\u0001\u0000\u0000\u0000$\u0005\u0001\u0000\u0000\u0000%#\u0001"+
		"\u0000\u0000\u0000&\'\u0005\u0003\u0000\u0000\'(\u0007\u0000\u0000\u0000"+
		"(\u0007\u0001\u0000\u0000\u0000)*\u0005\u0004\u0000\u0000*+\u0003\n\u0005"+
		"\u0000+\t\u0001\u0000\u0000\u0000,2\u0003\f\u0006\u0000-.\u0003\u0012"+
		"\t\u0000./\u0003\f\u0006\u0000/1\u0001\u0000\u0000\u00000-\u0001\u0000"+
		"\u0000\u000014\u0001\u0000\u0000\u000020\u0001\u0000\u0000\u000023\u0001"+
		"\u0000\u0000\u00003\u000b\u0001\u0000\u0000\u000042\u0001\u0000\u0000"+
		"\u00005:\u0003\u000e\u0007\u00006:\u0003\u0010\b\u000078\u0005\b\u0000"+
		"\u00008:\u0003\u000e\u0007\u000095\u0001\u0000\u0000\u000096\u0001\u0000"+
		"\u0000\u000097\u0001\u0000\u0000\u0000:\r\u0001\u0000\u0000\u0000;<\u0005"+
		"\u0001\u0000\u0000<=\u0005\f\u0000\u0000=>\u0005\u000e\u0000\u0000>?\u0005"+
		"\r\u0000\u0000?@\u0005\u000f\u0000\u0000@\u000f\u0001\u0000\u0000\u0000"+
		"AB\u0005\u0001\u0000\u0000BD\u0005\u0007\u0000\u0000CE\u0005\b\u0000\u0000"+
		"DC\u0001\u0000\u0000\u0000DE\u0001\u0000\u0000\u0000EF\u0001\u0000\u0000"+
		"\u0000FG\u0005\t\u0000\u0000G\u0011\u0001\u0000\u0000\u0000HI\u0007\u0001"+
		"\u0000\u0000I\u0013\u0001\u0000\u0000\u0000\u0006\u0017\u001c#29D";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}