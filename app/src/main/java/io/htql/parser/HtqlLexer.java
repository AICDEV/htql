// Generated from Htql.g4 by ANTLR 4.13.2
package io.htql.parser;
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class HtqlLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, SELECT=2, FROM=3, WHERE=4, AND=5, OR=6, IS=7, NOT=8, NULL=9, STAR=10, 
		COMMA=11, DOT=12, EQUALS=13, IDENTIFIER=14, STRING=15, WS=16, URL=17;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "SELECT", "FROM", "WHERE", "AND", "OR", "IS", "NOT", "NULL", 
			"STAR", "COMMA", "DOT", "EQUALS", "IDENTIFIER", "STRING", "WS", "URL"
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


	public HtqlLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Htql.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\u0011\u0097\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\n\u0001\n\u0001\u000b"+
		"\u0001\u000b\u0001\f\u0001\f\u0001\r\u0001\r\u0005\r^\b\r\n\r\f\ra\t\r"+
		"\u0001\u000e\u0001\u000e\u0005\u000ee\b\u000e\n\u000e\f\u000eh\t\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000f\u0004\u000fm\b\u000f\u000b\u000f"+
		"\f\u000fn\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0003\u0010|\b\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0004\u0010\u0083\b\u0010\u000b\u0010\f\u0010\u0084\u0001"+
		"\u0010\u0001\u0010\u0004\u0010\u0089\b\u0010\u000b\u0010\f\u0010\u008a"+
		"\u0003\u0010\u008d\b\u0010\u0001\u0010\u0001\u0010\u0005\u0010\u0091\b"+
		"\u0010\n\u0010\f\u0010\u0094\t\u0010\u0003\u0010\u0096\b\u0010\u0000\u0000"+
		"\u0011\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006"+
		"\r\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e"+
		"\u001d\u000f\u001f\u0010!\u0011\u0001\u0000\u0007\u0003\u0000AZ__az\u0004"+
		"\u000009AZ__az\u0003\u0000\n\n\r\r\'\'\u0003\u0000\t\n\r\r  \u0005\u0000"+
		"-.09AZ__az\u0001\u000009\b\u0000##%&.9==??AZ__az\u009f\u0000\u0001\u0001"+
		"\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001"+
		"\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000"+
		"\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000"+
		"\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000"+
		"\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000"+
		"\u0000\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000"+
		"\u0000\u0000\u001b\u0001\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000"+
		"\u0000\u0000\u001f\u0001\u0000\u0000\u0000\u0000!\u0001\u0000\u0000\u0000"+
		"\u0001#\u0001\u0000\u0000\u0000\u0003.\u0001\u0000\u0000\u0000\u00055"+
		"\u0001\u0000\u0000\u0000\u0007:\u0001\u0000\u0000\u0000\t@\u0001\u0000"+
		"\u0000\u0000\u000bD\u0001\u0000\u0000\u0000\rG\u0001\u0000\u0000\u0000"+
		"\u000fJ\u0001\u0000\u0000\u0000\u0011N\u0001\u0000\u0000\u0000\u0013S"+
		"\u0001\u0000\u0000\u0000\u0015U\u0001\u0000\u0000\u0000\u0017W\u0001\u0000"+
		"\u0000\u0000\u0019Y\u0001\u0000\u0000\u0000\u001b[\u0001\u0000\u0000\u0000"+
		"\u001db\u0001\u0000\u0000\u0000\u001fl\u0001\u0000\u0000\u0000!{\u0001"+
		"\u0000\u0000\u0000#$\u0005a\u0000\u0000$%\u0005t\u0000\u0000%&\u0005t"+
		"\u0000\u0000&\'\u0005r\u0000\u0000\'(\u0005i\u0000\u0000()\u0005b\u0000"+
		"\u0000)*\u0005u\u0000\u0000*+\u0005t\u0000\u0000+,\u0005e\u0000\u0000"+
		",-\u0005s\u0000\u0000-\u0002\u0001\u0000\u0000\u0000./\u0005S\u0000\u0000"+
		"/0\u0005E\u0000\u000001\u0005L\u0000\u000012\u0005E\u0000\u000023\u0005"+
		"C\u0000\u000034\u0005T\u0000\u00004\u0004\u0001\u0000\u0000\u000056\u0005"+
		"F\u0000\u000067\u0005R\u0000\u000078\u0005O\u0000\u000089\u0005M\u0000"+
		"\u00009\u0006\u0001\u0000\u0000\u0000:;\u0005W\u0000\u0000;<\u0005H\u0000"+
		"\u0000<=\u0005E\u0000\u0000=>\u0005R\u0000\u0000>?\u0005E\u0000\u0000"+
		"?\b\u0001\u0000\u0000\u0000@A\u0005A\u0000\u0000AB\u0005N\u0000\u0000"+
		"BC\u0005D\u0000\u0000C\n\u0001\u0000\u0000\u0000DE\u0005O\u0000\u0000"+
		"EF\u0005R\u0000\u0000F\f\u0001\u0000\u0000\u0000GH\u0005I\u0000\u0000"+
		"HI\u0005S\u0000\u0000I\u000e\u0001\u0000\u0000\u0000JK\u0005N\u0000\u0000"+
		"KL\u0005O\u0000\u0000LM\u0005T\u0000\u0000M\u0010\u0001\u0000\u0000\u0000"+
		"NO\u0005N\u0000\u0000OP\u0005U\u0000\u0000PQ\u0005L\u0000\u0000QR\u0005"+
		"L\u0000\u0000R\u0012\u0001\u0000\u0000\u0000ST\u0005*\u0000\u0000T\u0014"+
		"\u0001\u0000\u0000\u0000UV\u0005,\u0000\u0000V\u0016\u0001\u0000\u0000"+
		"\u0000WX\u0005.\u0000\u0000X\u0018\u0001\u0000\u0000\u0000YZ\u0005=\u0000"+
		"\u0000Z\u001a\u0001\u0000\u0000\u0000[_\u0007\u0000\u0000\u0000\\^\u0007"+
		"\u0001\u0000\u0000]\\\u0001\u0000\u0000\u0000^a\u0001\u0000\u0000\u0000"+
		"_]\u0001\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000`\u001c\u0001\u0000"+
		"\u0000\u0000a_\u0001\u0000\u0000\u0000bf\u0005\'\u0000\u0000ce\b\u0002"+
		"\u0000\u0000dc\u0001\u0000\u0000\u0000eh\u0001\u0000\u0000\u0000fd\u0001"+
		"\u0000\u0000\u0000fg\u0001\u0000\u0000\u0000gi\u0001\u0000\u0000\u0000"+
		"hf\u0001\u0000\u0000\u0000ij\u0005\'\u0000\u0000j\u001e\u0001\u0000\u0000"+
		"\u0000km\u0007\u0003\u0000\u0000lk\u0001\u0000\u0000\u0000mn\u0001\u0000"+
		"\u0000\u0000nl\u0001\u0000\u0000\u0000no\u0001\u0000\u0000\u0000op\u0001"+
		"\u0000\u0000\u0000pq\u0006\u000f\u0000\u0000q \u0001\u0000\u0000\u0000"+
		"rs\u0005h\u0000\u0000st\u0005t\u0000\u0000tu\u0005t\u0000\u0000u|\u0005"+
		"p\u0000\u0000vw\u0005h\u0000\u0000wx\u0005t\u0000\u0000xy\u0005t\u0000"+
		"\u0000yz\u0005p\u0000\u0000z|\u0005s\u0000\u0000{r\u0001\u0000\u0000\u0000"+
		"{v\u0001\u0000\u0000\u0000|}\u0001\u0000\u0000\u0000}~\u0005:\u0000\u0000"+
		"~\u007f\u0005/\u0000\u0000\u007f\u0080\u0005/\u0000\u0000\u0080\u0082"+
		"\u0001\u0000\u0000\u0000\u0081\u0083\u0007\u0004\u0000\u0000\u0082\u0081"+
		"\u0001\u0000\u0000\u0000\u0083\u0084\u0001\u0000\u0000\u0000\u0084\u0082"+
		"\u0001\u0000\u0000\u0000\u0084\u0085\u0001\u0000\u0000\u0000\u0085\u008c"+
		"\u0001\u0000\u0000\u0000\u0086\u0088\u0005:\u0000\u0000\u0087\u0089\u0007"+
		"\u0005\u0000\u0000\u0088\u0087\u0001\u0000\u0000\u0000\u0089\u008a\u0001"+
		"\u0000\u0000\u0000\u008a\u0088\u0001\u0000\u0000\u0000\u008a\u008b\u0001"+
		"\u0000\u0000\u0000\u008b\u008d\u0001\u0000\u0000\u0000\u008c\u0086\u0001"+
		"\u0000\u0000\u0000\u008c\u008d\u0001\u0000\u0000\u0000\u008d\u0095\u0001"+
		"\u0000\u0000\u0000\u008e\u0092\u0005/\u0000\u0000\u008f\u0091\u0007\u0006"+
		"\u0000\u0000\u0090\u008f\u0001\u0000\u0000\u0000\u0091\u0094\u0001\u0000"+
		"\u0000\u0000\u0092\u0090\u0001\u0000\u0000\u0000\u0092\u0093\u0001\u0000"+
		"\u0000\u0000\u0093\u0096\u0001\u0000\u0000\u0000\u0094\u0092\u0001\u0000"+
		"\u0000\u0000\u0095\u008e\u0001\u0000\u0000\u0000\u0095\u0096\u0001\u0000"+
		"\u0000\u0000\u0096\"\u0001\u0000\u0000\u0000\n\u0000_fn{\u0084\u008a\u008c"+
		"\u0092\u0095\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}