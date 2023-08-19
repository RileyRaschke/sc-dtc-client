
// Data and Trading Communications Protocol (DTC Protocol)

// This protocol is in the public domain and freely usable by anyone.

// Documentation: http://DTCprotocol.org/index.php?page=doc_DTCMessageDocumentation.php

#pragma once

#include <cstdint>

namespace DTC
{
// In general because these structures are placed into a stream of data 
// and the beginning position of the structure in the stream is variable 
// and not to any particular boundary, structure member alignment here is 
// not beneficial. So this could be changed to #pragma pack(1). It is 
// maintained at 8 since there is a reliance on that alignment for existing
// compiled code.
#pragma pack(push, 8)

	// DTC protocol version
	const int32_t CURRENT_VERSION = 8;

	// Text string lengths when using fixed length string binary encoding. 
	const int32_t USERNAME_PASSWORD_LENGTH = 32;
	const int32_t SYMBOL_EXCHANGE_DELIMITER_LENGTH = 4;
	const int32_t SYMBOL_LENGTH = 64;
	const int32_t EXCHANGE_LENGTH = 16;
	const int32_t UNDERLYING_SYMBOL_LENGTH = 32;
	const int32_t SYMBOL_DESCRIPTION_LENGTH = 64;
	const int32_t EXCHANGE_DESCRIPTION_LENGTH = 48;
	const int32_t ORDER_ID_LENGTH = 32;
	const int32_t TRADE_ACCOUNT_LENGTH = 32;
	const int32_t TEXT_DESCRIPTION_LENGTH = 96;
	const int32_t TEXT_MESSAGE_LENGTH = 256;
	const int32_t ORDER_FREE_FORM_TEXT_LENGTH = 48;
	const int32_t CLIENT_SERVER_NAME_LENGTH = 48;
	const int32_t GENERAL_IDENTIFIER_LENGTH = 64;
	const int32_t CURRENCY_CODE_LENGTH = 8;
	const int32_t ORDER_FILL_EXECUTION_LENGTH = 64;
	const int32_t PRICE_STRING_LENGTH = 16;

	//----Message types----

	// Authentication and connection monitoring
	const uint16_t LOGON_REQUEST = 1;
	const uint16_t LOGON_RESPONSE = 2;
	const uint16_t HEARTBEAT = 3;
	const uint16_t LOGOFF = 5;
	const uint16_t ENCODING_REQUEST = 6;
	const uint16_t ENCODING_RESPONSE = 7;

	// Market data
	const uint16_t MARKET_DATA_REQUEST = 101;
	const uint16_t MARKET_DATA_REJECT = 103;
	const uint16_t MARKET_DATA_SNAPSHOT = 104;
	//125

	const uint16_t MARKET_DATA_UPDATE_TRADE = 107;
	const uint16_t MARKET_DATA_UPDATE_TRADE_COMPACT = 112;
	//126
	const uint16_t MARKET_DATA_UPDATE_LAST_TRADE_SNAPSHOT = 134;
	const uint16_t MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR = 137;
	const uint16_t MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_2 = 146;
	const uint16_t MARKET_DATA_UPDATE_TRADE_NO_TIMESTAMP = 142;

	const uint16_t MARKET_DATA_UPDATE_BID_ASK = 108;
	const uint16_t MARKET_DATA_UPDATE_BID_ASK_COMPACT = 117;
	//127
	const uint16_t MARKET_DATA_UPDATE_BID_ASK_NO_TIMESTAMP = 143;
	const uint16_t MARKET_DATA_UPDATE_BID_ASK_FLOAT_WITH_MICROSECONDS = 144;

	const uint16_t MARKET_DATA_UPDATE_SESSION_OPEN = 120;
	//128
	const uint16_t MARKET_DATA_UPDATE_SESSION_HIGH = 114;
	//129
	const uint16_t MARKET_DATA_UPDATE_SESSION_LOW = 115;
	//130
	const uint16_t MARKET_DATA_UPDATE_SESSION_VOLUME = 113;
	const uint16_t MARKET_DATA_UPDATE_OPEN_INTEREST = 124;
	const uint16_t MARKET_DATA_UPDATE_SESSION_SETTLEMENT = 119;
	//131
	const uint16_t MARKET_DATA_UPDATE_SESSION_NUM_TRADES = 135;
	const uint16_t MARKET_DATA_UPDATE_TRADING_SESSION_DATE = 136;

	const uint16_t MARKET_DEPTH_REQUEST = 102;
	const uint16_t MARKET_DEPTH_REJECT = 121;
	const uint16_t MARKET_DEPTH_SNAPSHOT_LEVEL = 122;
	//132
	const uint16_t MARKET_DEPTH_SNAPSHOT_LEVEL_FLOAT = 145;
	const uint16_t MARKET_DEPTH_UPDATE_LEVEL = 106;
	const uint16_t MARKET_DEPTH_UPDATE_LEVEL_FLOAT_WITH_MILLISECONDS = 140;
	const uint16_t MARKET_DEPTH_UPDATE_LEVEL_NO_TIMESTAMP = 141;
	//133

	const uint16_t MARKET_DATA_FEED_STATUS = 100;
	const uint16_t MARKET_DATA_FEED_SYMBOL_STATUS = 116;
	const uint16_t TRADING_SYMBOL_STATUS = 138;

	const uint16_t MARKET_ORDERS_REQUEST = 150;
	const uint16_t MARKET_ORDERS_REJECT = 151;

	const uint16_t MARKET_ORDERS_ADD = 152;
	const uint16_t MARKET_ORDERS_MODIFY = 153;
	const uint16_t MARKET_ORDERS_REMOVE = 154;
	const uint16_t MARKET_ORDERS_SNAPSHOT_MESSAGE_BOUNDARY = 155;


	// Order entry and modification
	const uint16_t SUBMIT_NEW_SINGLE_ORDER = 208;
	//206

	const uint16_t SUBMIT_NEW_OCO_ORDER = 201;
	//207
	const uint16_t SUBMIT_FLATTEN_POSITION_ORDER = 209;
	const uint16_t FLATTEN_POSITIONS_FOR_TRADE_ACCOUNT = 210;

	const uint16_t CANCEL_ORDER = 203;

	const uint16_t CANCEL_REPLACE_ORDER = 204;
	//205

	// Trading related
	const uint16_t OPEN_ORDERS_REQUEST = 300;
	const uint16_t OPEN_ORDERS_REJECT = 302;

	const uint16_t ORDER_UPDATE = 301;

	const uint16_t HISTORICAL_ORDER_FILLS_REQUEST = 303;
	const uint16_t HISTORICAL_ORDER_FILLS_REJECT = 308;
	const uint16_t HISTORICAL_ORDER_FILL_RESPONSE = 304;

	const uint16_t CURRENT_POSITIONS_REQUEST = 305;
	const uint16_t CURRENT_POSITIONS_REJECT = 307;
	const uint16_t POSITION_UPDATE = 306;
	const uint16_t ADD_CORRECTING_ORDER_FILL = 309;
	const uint16_t CORRECTING_ORDER_FILL_RESPONSE = 310;

	// Account list
	const uint16_t TRADE_ACCOUNTS_REQUEST = 400;
	const uint16_t TRADE_ACCOUNT_RESPONSE = 401;

	// Symbol discovery and security definitions
	const uint16_t EXCHANGE_LIST_REQUEST = 500;
	const uint16_t EXCHANGE_LIST_RESPONSE = 501;

	const uint16_t SYMBOLS_FOR_EXCHANGE_REQUEST = 502;
	const uint16_t UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST = 503;
	const uint16_t SYMBOLS_FOR_UNDERLYING_REQUEST = 504;
	const uint16_t SECURITY_DEFINITION_FOR_SYMBOL_REQUEST = 506;
	const uint16_t SECURITY_DEFINITION_RESPONSE = 507;

	const uint16_t SYMBOL_SEARCH_REQUEST = 508;

	const uint16_t SECURITY_DEFINITION_REJECT = 509;

	// Account Balance Data
	const uint16_t ACCOUNT_BALANCE_REQUEST = 601;
	const uint16_t ACCOUNT_BALANCE_REJECT = 602;	
	const uint16_t ACCOUNT_BALANCE_UPDATE = 600;
	const uint16_t ACCOUNT_BALANCE_ADJUSTMENT = 607;
	const uint16_t ACCOUNT_BALANCE_ADJUSTMENT_REJECT = 608;
	const uint16_t ACCOUNT_BALANCE_ADJUSTMENT_COMPLETE = 609;
	const uint16_t HISTORICAL_ACCOUNT_BALANCES_REQUEST = 603;
	const uint16_t HISTORICAL_ACCOUNT_BALANCES_REJECT = 604;
	const uint16_t HISTORICAL_ACCOUNT_BALANCE_RESPONSE = 606;

	// Logging
	const uint16_t USER_MESSAGE = 700;
	const uint16_t GENERAL_LOG_MESSAGE = 701;
	const uint16_t ALERT_MESSAGE = 702;

	const uint16_t JOURNAL_ENTRY_ADD = 703;
	const uint16_t JOURNAL_ENTRIES_REQUEST = 704;
	const uint16_t JOURNAL_ENTRIES_REJECT = 705;
	const uint16_t JOURNAL_ENTRY_RESPONSE = 706;
	
	// Historical price data
	const uint16_t HISTORICAL_PRICE_DATA_REQUEST = 800;
	const uint16_t HISTORICAL_PRICE_DATA_RESPONSE_HEADER = 801;
	const uint16_t HISTORICAL_PRICE_DATA_REJECT = 802;
	const uint16_t HISTORICAL_PRICE_DATA_RECORD_RESPONSE = 803;
	const uint16_t HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE = 804;
	//805
	//806
	const uint16_t HISTORICAL_PRICE_DATA_RESPONSE_TRAILER = 807;

	// Historical market depth data
	const uint16_t HISTORICAL_MARKET_DEPTH_DATA_REQUEST = 900;
	const uint16_t HISTORICAL_MARKET_DEPTH_DATA_RESPONSE_HEADER = 901;
	const uint16_t HISTORICAL_MARKET_DEPTH_DATA_REJECT = 902;
	const uint16_t HISTORICAL_MARKET_DEPTH_DATA_RECORD_RESPONSE = 903;

	/*==========================================================================*/
	// Standard UNIX date-time value. In seconds.
	typedef int64_t t_DateTime;

	// This is a 32 bit UNIX date-time value used in messages where compactness is an important consideration. Or, where only the Date is needed. In seconds.
	typedef uint32_t t_DateTime4Byte;

	// UNIX date-time value with fractional portion for milliseconds.
	typedef double t_DateTimeWithMilliseconds;

	// Standard UNIX date-time value in milliseconds.
	typedef int64_t t_DateTimeWithMillisecondsInt;

	// Standard UNIX date-time value in microseconds.
	typedef int64_t t_DateTimeWithMicrosecondsInt;

	inline int64_t DateTimeWithMillisecondsToDateTimeWithMicrosecondsInt(const t_DateTimeWithMilliseconds DateTime)
	{
		return static_cast <int64_t>(DateTime * 1000000.0 + 0.5);
	}

	/*==========================================================================*/
	enum EncodingEnum : int32_t
	{ BINARY_ENCODING = 0
	, BINARY_WITH_VARIABLE_LENGTH_STRINGS = 1
	, JSON_ENCODING = 2
	, JSON_COMPACT_ENCODING = 3
	, PROTOCOL_BUFFERS = 4
	};

	inline const char* GetEncodingName(EncodingEnum Encoding)
	{
		switch (Encoding)
		{
			case BINARY_ENCODING:
			return "Binary";

			case BINARY_WITH_VARIABLE_LENGTH_STRINGS:
			return "Binary VLS";

			case JSON_ENCODING:
			return "JSON";

			case JSON_COMPACT_ENCODING:
			return "JSON Compact";

			case PROTOCOL_BUFFERS:
			return "Protocol Buffers";

			default:
			return "Unknown";
		}
	}

	/*==========================================================================*/
	enum LogonStatusEnum : int32_t
	{ LOGON_SUCCESS = 1
	, LOGON_ERROR = 2
	, LOGON_ERROR_NO_RECONNECT = 3
	, LOGON_RECONNECT_NEW_ADDRESS = 4
	};

	/*==========================================================================*/
	enum RequestActionEnum : int32_t
	{ SUBSCRIBE = 1
	, UNSUBSCRIBE = 2
	, SNAPSHOT = 3
	, SNAPSHOT_WITH_INTERVAL_UPDATES = 4
	};

	/*==========================================================================*/
	enum UnbundledTradeIndicatorEnum : int8_t
	{ UNBUNDLED_TRADE_NONE = 0
	, FIRST_SUB_TRADE_OF_UNBUNDLED_TRADE = 1
	, LAST_SUB_TRADE_OF_UNBUNDLED_TRADE = 2
	};

	/*==========================================================================*/
	enum OrderStatusEnum : int32_t
	{ ORDER_STATUS_UNSPECIFIED = 0
	, ORDER_STATUS_ORDER_SENT = 1
	, ORDER_STATUS_PENDING_OPEN = 2
	, ORDER_STATUS_PENDING_CHILD = 3
	, ORDER_STATUS_OPEN = 4
	, ORDER_STATUS_PENDING_CANCEL_REPLACE = 5
	, ORDER_STATUS_PENDING_CANCEL = 6
	, ORDER_STATUS_FILLED = 7
	, ORDER_STATUS_CANCELED = 8
	, ORDER_STATUS_REJECTED = 9
	, ORDER_STATUS_PARTIALLY_FILLED = 10
	};

	/*==========================================================================*/
	enum OrderUpdateReasonEnum : int32_t
	{ ORDER_UPDATE_REASON_UNSET = 0
	, OPEN_ORDERS_REQUEST_RESPONSE = 1
	, NEW_ORDER_ACCEPTED = 2
	, GENERAL_ORDER_UPDATE = 3
	, ORDER_FILLED = 4
	, ORDER_FILLED_PARTIALLY = 5
	, ORDER_CANCELED = 6
	, ORDER_CANCEL_REPLACE_COMPLETE = 7
	, NEW_ORDER_REJECTED = 8
	, ORDER_CANCEL_REJECTED = 9
	, ORDER_CANCEL_REPLACE_REJECTED = 10
	};
	
	/*==========================================================================*/
	enum AtBidOrAskEnum8 : uint8_t
	{ BID_ASK_UNSET_8 = 0
	, AT_BID_8 = 1
	, AT_ASK_8 = 2
	};

	/*==========================================================================*/
	enum AtBidOrAskEnum : uint16_t
	{ BID_ASK_UNSET = 0
	, AT_BID = 1
	, AT_ASK = 2
	};

	/*==========================================================================*/
	enum MarketDepthUpdateTypeEnum : uint8_t
	{ MARKET_DEPTH_UNSET = 0
	, MARKET_DEPTH_INSERT_UPDATE_LEVEL = 1 // Insert or update depth at the given price level
	, MARKET_DEPTH_DELETE_LEVEL = 2 // Delete depth at the given price level
	};

	/*==========================================================================*/
	enum FinalUpdateInBatchEnum : uint8_t
	{
		FINAL_UPDATE_UNSET = 0
		, FINAL_UPDATE_TRUE = 1
		, FINAL_UPDATE_FALSE = 2
		, FINAL_UPDATE_BEGIN_BATCH = 3
	};

	/*==========================================================================*/
	enum MessageSetBoundaryEnum : uint8_t
	{
		MESSAGE_SET_BOUNDARY_UNSET = 0
		, MESSAGE_SET_BOUNDARY_BEGIN = 1
		, MESSAGE_SET_BOUNDARY_END = 2
	};
	
	/*==========================================================================*/
	enum OrderTypeEnum : int32_t
	{ ORDER_TYPE_UNSET = 0
	, ORDER_TYPE_MARKET = 1
	, ORDER_TYPE_LIMIT = 2
	, ORDER_TYPE_STOP = 3
	, ORDER_TYPE_STOP_LIMIT = 4
	, ORDER_TYPE_MARKET_IF_TOUCHED = 5
	, ORDER_TYPE_LIMIT_IF_TOUCHED = 6
	, ORDER_TYPE_MARKET_LIMIT = 7
	};
	
	/*==========================================================================*/
	enum TimeInForceEnum : int32_t
	{ TIF_UNSET = 0
	, TIF_DAY = 1
	, TIF_GOOD_TILL_CANCELED = 2
	, TIF_GOOD_TILL_DATE_TIME = 3
	, TIF_IMMEDIATE_OR_CANCEL = 4
	, TIF_ALL_OR_NONE = 5
	, TIF_FILL_OR_KILL = 6
	};
	
	/*==========================================================================*/
	enum BuySellEnum : int32_t
	{ BUY_SELL_UNSET = 0
	, BUY = 1
	, SELL = 2
	};

	/*==========================================================================*/
	enum OpenCloseTradeEnum : int32_t
	{ TRADE_UNSET = 0
	, TRADE_OPEN = 1
	, TRADE_CLOSE = 2
	};

	/*==========================================================================*/
	enum PartialFillHandlingEnum : int8_t
	{ PARTIAL_FILL_UNSET = 0
	, PARTIAL_FILL_HANDLING_REDUCE_QUANTITY = 1
	, PARTIAL_FILL_HANDLING_IMMEDIATE_CANCEL = 2
	};

	/*==========================================================================*/
	enum MarketDataFeedStatusEnum : int32_t
	{ MARKET_DATA_FEED_STATUS_UNSET = 0
	, MARKET_DATA_FEED_UNAVAILABLE = 1
	, MARKET_DATA_FEED_AVAILABLE = 2
	};

	/*==========================================================================*/
	enum TradingStatusEnum : int8_t
	{ TRADING_STATUS_UNKNOWN = 0
	, TRADING_STATUS_PRE_OPEN = 1
	, TRADING_STATUS_OPEN = 2
	, TRADING_STATUS_CLOSE = 3
	, TRADING_STATUS_TRADING_HALT = 4
	};

	inline const char* GetTradingStatusString(TradingStatusEnum TradingStatus)
	{
		switch (TradingStatus)
		{

		case TRADING_STATUS_PRE_OPEN:
			return "Pre-open";

		case TRADING_STATUS_OPEN:
			return "Open";

		case TRADING_STATUS_CLOSE:
			return "Closed";

		case TRADING_STATUS_TRADING_HALT:
			return "Trading Halt";

		case TRADING_STATUS_UNKNOWN:
		default:
			return "Unknown";
		}
	}


	/*==========================================================================*/
	enum PriceDisplayFormatEnum : int32_t
	{ PRICE_DISPLAY_FORMAT_UNSET =  -1
	//The following formats indicate the number of decimal places to be displayed
	, PRICE_DISPLAY_FORMAT_DECIMAL_0 = 0
	, PRICE_DISPLAY_FORMAT_DECIMAL_1 = 1
	, PRICE_DISPLAY_FORMAT_DECIMAL_2 = 2
	, PRICE_DISPLAY_FORMAT_DECIMAL_3 = 3
	, PRICE_DISPLAY_FORMAT_DECIMAL_4 = 4
	, PRICE_DISPLAY_FORMAT_DECIMAL_5 = 5
	, PRICE_DISPLAY_FORMAT_DECIMAL_6 = 6
	, PRICE_DISPLAY_FORMAT_DECIMAL_7 = 7
	, PRICE_DISPLAY_FORMAT_DECIMAL_8 = 8
	, PRICE_DISPLAY_FORMAT_DECIMAL_9 = 9
	//The following formats are fractional formats
	, PRICE_DISPLAY_FORMAT_DENOMINATOR_256 = 356
	, PRICE_DISPLAY_FORMAT_DENOMINATOR_128 = 228
	, PRICE_DISPLAY_FORMAT_DENOMINATOR_64 = 164
	, PRICE_DISPLAY_FORMAT_DENOMINATOR_32_EIGHTHS = 140
	, PRICE_DISPLAY_FORMAT_DENOMINATOR_32_QUARTERS = 136
	, PRICE_DISPLAY_FORMAT_DENOMINATOR_32_HALVES = 134
	, PRICE_DISPLAY_FORMAT_DENOMINATOR_32 = 132 
	, PRICE_DISPLAY_FORMAT_DENOMINATOR_16 = 116
	, PRICE_DISPLAY_FORMAT_DENOMINATOR_8 = 108
	, PRICE_DISPLAY_FORMAT_DENOMINATOR_4 = 104
	, PRICE_DISPLAY_FORMAT_DENOMINATOR_2 = 102
	};

	/*==========================================================================*/
	enum SecurityTypeEnum : int32_t
	{ SECURITY_TYPE_UNSET = 0
	, SECURITY_TYPE_FUTURES = 1
	, SECURITY_TYPE_STOCK = 2
	, SECURITY_TYPE_FOREX = 3 // CryptoCurrencies also go into this category
	, SECURITY_TYPE_INDEX = 4
	, SECURITY_TYPE_FUTURES_STRATEGY = 5
	, SECURITY_TYPE_FUTURES_OPTION = 7
	, SECURITY_TYPE_STOCK_OPTION = 6
	, SECURITY_TYPE_INDEX_OPTION = 8
	, SECURITY_TYPE_BOND = 9
	, SECURITY_TYPE_MUTUAL_FUND = 10
	};

	enum PutCallEnum : uint8_t
	{ PC_UNSET = 0
	, PC_CALL = 1
	, PC_PUT = 2
	};

	enum SearchTypeEnum : int32_t
	{ SEARCH_TYPE_UNSET = 0
	, SEARCH_TYPE_BY_SYMBOL = 1
	, SEARCH_TYPE_BY_DESCRIPTION = 2
	};
	
	/*==========================================================================*/
	enum HistoricalDataIntervalEnum : int32_t
	{ INTERVAL_TICK = 0
	, INTERVAL_1_SECOND = 1
	, INTERVAL_2_SECONDS = 2
	, INTERVAL_4_SECONDS = 4
	, INTERVAL_5_SECONDS = 5
	, INTERVAL_10_SECONDS = 10
	, INTERVAL_30_SECONDS = 30
	, INTERVAL_1_MINUTE = 60
	, INTERVAL_5_MINUTE = 300
	, INTERVAL_10_MINUTE = 600
	, INTERVAL_15_MINUTE = 900
	, INTERVAL_30_MINUTE = 1800
	, INTERVAL_1_HOUR = 3600
	, INTERVAL_2_HOURS = 7200
	, INTERVAL_1_DAY = 86400
	, INTERVAL_1_WEEK = 604800
	};

	enum HistoricalPriceDataRejectReasonCodeEnum : int16_t
	{ HPDR_UNSET = 0
	, HPDR_UNABLE_TO_SERVE_DATA_RETRY_IN_SPECIFIED_SECONDS = 1
	, HPDR_UNABLE_TO_SERVE_DATA_DO_NOT_RETRY = 2
	, HPDR_DATA_REQUEST_OUTSIDE_BOUNDS_OF_AVAILABLE_DATA = 3
	, HPDR_GENERAL_REJECT_ERROR = 4
	};

	/*==========================================================================*/
	enum TradeConditionEnum : int8_t
	{
		TRADE_CONDITION_NONE = 0
		, TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE = 1
		, TRADE_CONDITION_ODD_LOT_EQUITY_TRADE = 2
	};

/*****************************************************************************
// s_MessageHeader

For binary encodings, every DTC message starts with these these two fields,
which can be used to determine how to handle the rest of the message.
----------------------------------------------------------------------------*/
struct s_MessageHeader
{
	uint16_t Size = 0;
	uint16_t Type = 0;
};

	/*==========================================================================*/
	struct s_EncodingRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = ENCODING_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;
		int32_t ProtocolVersion = CURRENT_VERSION;
		EncodingEnum Encoding = BINARY_ENCODING;
		char ProtocolType[4] = "DTC";

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);		

		int32_t GetProtocolVersion() const;
		EncodingEnum GetEncoding() const;
		const char* GetProtocolType();
		void SetProtocolType(const char* NewValue);
	};

	/*==========================================================================*/
	struct s_EncodingResponse
	{
		static constexpr uint16_t MESSAGE_TYPE = ENCODING_RESPONSE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;
		int32_t ProtocolVersion = CURRENT_VERSION;
		EncodingEnum Encoding = BINARY_ENCODING;

		char ProtocolType[4] = "DTC";

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetProtocolVersion() const;
		EncodingEnum GetEncoding() const;
		const char* GetProtocolType();
		void SetProtocolType(const char* NewValue);
	};

	/*==========================================================================*/
	struct s_LogonRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = LOGON_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t ProtocolVersion = CURRENT_VERSION;
		char Username[USERNAME_PASSWORD_LENGTH] = {};
		char Password[USERNAME_PASSWORD_LENGTH] = {};
		char GeneralTextData[GENERAL_IDENTIFIER_LENGTH] = {};
		int32_t Integer_1 = 0;
		int32_t Integer_2 = 0;
		int32_t HeartbeatIntervalInSeconds = 0;
		int32_t Unused1 = 0;
		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};
		char HardwareIdentifier[GENERAL_IDENTIFIER_LENGTH] = {};
		char ClientName[32] = {};
		int32_t MarketDataTransmissionInterval = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetProtocolVersion() const;
		const char* GetUsername();
		void SetUsername(const char* NewValue);
		const char* GetPassword();
		void SetPassword(const char* NewValue);
		const char* GetGeneralTextData();
		void SetGeneralTextData(const char* NewValue);
		int32_t GetInteger_1() const;
		int32_t GetInteger_2() const;
		int32_t GetHeartbeatIntervalInSeconds() const;
		const char* GetTradeAccount();
		void SetTradeAccount(const char* NewValue);
		const char* GetHardwareIdentifier();
		void SetHardwareIdentifier(const char* NewValue);
		const char* GetClientName();
		void SetClientName(const char* NewValue);
		int32_t GetMarketDataTransmissionInterval() const;
	};

	/*==========================================================================*/
	struct s_LogonResponse
	{
		static constexpr uint16_t MESSAGE_TYPE = LOGON_RESPONSE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t ProtocolVersion = CURRENT_VERSION;
		LogonStatusEnum Result = LOGON_SUCCESS;
		char ResultText[TEXT_DESCRIPTION_LENGTH] = {};
		char ReconnectAddress[64] = {};
		int32_t Integer_1 = 0;
		char ServerName[60] = {};
		uint8_t MarketDepthUpdatesBestBidAndAsk = 0;
		uint8_t TradingIsSupported = 0;
		uint8_t OCOOrdersSupported = 0;
		uint8_t OrderCancelReplaceSupported = 1;
		char SymbolExchangeDelimiter[SYMBOL_EXCHANGE_DELIMITER_LENGTH] = {};
		uint8_t SecurityDefinitionsSupported = 0;
		uint8_t HistoricalPriceDataSupported = 0;
		uint8_t ResubscribeWhenMarketDataFeedAvailable = 0;
		uint8_t MarketDepthIsSupported = 1;
		uint8_t OneHistoricalPriceDataRequestPerConnection = 0;
		uint8_t BracketOrdersSupported = 0;
		uint8_t Unused_1 = 0;
		uint8_t UsesMultiplePositionsPerSymbolAndTradeAccount = 0;
		uint8_t MarketDataSupported = 1;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetProtocolVersion() const;
		LogonStatusEnum GetResult() const;
		const char* GetResultText();
		void SetResultText(const char* NewValue);
		const char* GetReconnectAddress();
		void SetReconnectAddress(const char* NewValue);
		int32_t GetInteger_1() const;
		const char* GetServerName();
		void SetServerName(const char* NewValue);		
		uint8_t GetMarketDepthUpdatesBestBidAndAsk() const;
		uint8_t GetTradingIsSupported() const;
		uint8_t GetOCOOrdersSupported() const;
		uint8_t GetOrderCancelReplaceSupported() const;
		const char* GetSymbolExchangeDelimiter();
		void SetSymbolExchangeDelimiter(const char* NewValue);
		uint8_t GetSecurityDefinitionsSupported() const;
		uint8_t GetHistoricalPriceDataSupported() const;
		uint8_t GetResubscribeWhenMarketDataFeedAvailable() const;
		uint8_t GetMarketDepthIsSupported() const;
		uint8_t GetOneHistoricalPriceDataRequestPerConnection() const;
		uint8_t GetBracketOrdersSupported() const;
		uint8_t GetUsesMultiplePositionsPerSymbolAndTradeAccount() const;
		uint8_t GetMarketDataSupported() const;
	};

	/*==========================================================================*/
	struct s_Logoff
	{
		static constexpr uint16_t MESSAGE_TYPE = LOGOFF;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;
		char Reason[TEXT_DESCRIPTION_LENGTH] = {};
		uint8_t DoNotReconnect = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		const char* GetReason();
		void SetReason(const char* NewValue);
		uint8_t GetDoNotReconnect() const;

	};

	/*==========================================================================*/
	struct s_Heartbeat
	{
		static constexpr uint16_t MESSAGE_TYPE = HEARTBEAT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t NumDroppedMessages = 0;
		t_DateTime CurrentDateTime = 0;
				
		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetNumDroppedMessages() const;
		t_DateTime GetCurrentDateTime() const;
	};

	/*==========================================================================*/
	struct s_MarketDataFeedStatus
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_FEED_STATUS;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		MarketDataFeedStatusEnum Status = MARKET_DATA_FEED_STATUS_UNSET;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		MarketDataFeedStatusEnum GetStatus() const;
	};

	/*==========================================================================*/
	struct s_MarketDataFeedSymbolStatus
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_FEED_SYMBOL_STATUS;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;
		uint32_t SymbolID = 0;
		MarketDataFeedStatusEnum Status = MARKET_DATA_FEED_STATUS_UNSET;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		MarketDataFeedStatusEnum GetStatus() const;
	};

	/*==========================================================================*/
	struct s_TradingSymbolStatus
	{
		static constexpr uint16_t MESSAGE_TYPE = TRADING_SYMBOL_STATUS;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;
		uint32_t SymbolID = 0;
		TradingStatusEnum Status = TRADING_STATUS_UNKNOWN;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		TradingStatusEnum GetStatus() const;
	};

	/*==========================================================================*/
	struct s_MarketDataRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		RequestActionEnum RequestAction = SUBSCRIBE;
		uint32_t SymbolID = 0;
		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};
		uint32_t IntervalForSnapshotUpdatesInMilliseconds = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		RequestActionEnum GetRequestAction() const;
		uint32_t GetSymbolID() const;
		const char* GetSymbol();
		void SetSymbol(const char* NewValue);
		const char* GetExchange();
		void SetExchange(const char* NewValue);
		uint32_t GetIntervalForSnapshotUpdatesInMilliseconds() const;
	};

	/*==========================================================================*/
	struct s_MarketDepthRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DEPTH_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		RequestActionEnum RequestAction = SUBSCRIBE;
		uint32_t SymbolID = 0;
		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};
		int32_t NumLevels = 0;

		
		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		RequestActionEnum GetRequestAction() const;
		uint32_t GetSymbolID() const;
		const char* GetSymbol();
		void SetSymbol(const char* NewValue);
		const char* GetExchange();
		void SetExchange(const char* NewValue);
		int32_t GetNumLevels() const;
	};

	/*==========================================================================*/
	struct s_MarketDataReject
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_REJECT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;
		uint32_t SymbolID = 0;
		char RejectText[TEXT_DESCRIPTION_LENGTH] = {};
				
		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		const char* GetRejectText();
		void SetRejectText(const char* NewValue);
	};

	/*==========================================================================*/
	struct s_MarketDataSnapshot
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_SNAPSHOT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		double SessionSettlementPrice = DBL_MAX;
		double SessionOpenPrice = DBL_MAX;
		double SessionHighPrice = DBL_MAX;
		double SessionLowPrice = DBL_MAX;
		double SessionVolume = DBL_MAX;
		uint32_t SessionNumTrades = UINT_MAX;
		uint32_t OpenInterest = UINT_MAX;

		double BidPrice = DBL_MAX;
		double AskPrice = DBL_MAX;
		double AskQuantity = DBL_MAX;
		double BidQuantity = DBL_MAX;
		double LastTradePrice = DBL_MAX;
		double LastTradeVolume = DBL_MAX;
		t_DateTimeWithMilliseconds LastTradeDateTime = 0;
		t_DateTimeWithMilliseconds BidAskDateTime = 0;
		t_DateTime4Byte SessionSettlementDateTime = 0;
		t_DateTime4Byte TradingSessionDate = 0;
		TradingStatusEnum TradingStatus = TRADING_STATUS_UNKNOWN;
		t_DateTimeWithMilliseconds MarketDepthUpdateDateTime = 0;
				
		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		double GetSessionSettlementPrice() const;
		double GetSessionOpenPrice() const;
		double GetSessionHighPrice() const;
		double GetSessionLowPrice() const;
		double GetSessionVolume() const;
		uint32_t GetSessionNumTrades() const;
		uint32_t GetOpenInterest() const;
		double GetBidPrice() const;
		double GetAskPrice() const;
		double GetAskQuantity() const;
		double GetBidQuantity() const;
		double GetLastTradePrice() const;
		double GetLastTradeVolume() const;
		t_DateTimeWithMilliseconds GetLastTradeDateTime() const;
		t_DateTimeWithMilliseconds GetBidAskDateTime() const;
		t_DateTime4Byte GetSessionSettlementDateTime() const;
		t_DateTime4Byte GetTradingSessionDate() const;
		TradingStatusEnum GetTradingStatus() const;
		t_DateTimeWithMilliseconds GetMarketDepthUpdateDateTime() const;
	};

	/*==========================================================================*/
	struct s_MarketDepthSnapshotLevel
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DEPTH_SNAPSHOT_LEVEL;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		AtBidOrAskEnum Side = BID_ASK_UNSET;
		double Price = 0;
		double Quantity = 0;
		uint16_t  Level = 0;

		uint8_t IsFirstMessageInBatch = 0;
		uint8_t IsLastMessageInBatch = 0;

		t_DateTimeWithMilliseconds DateTime = 0;

		uint32_t NumOrders = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		AtBidOrAskEnum GetSide() const;
		double GetPrice() const;
		double GetQuantity() const;
		uint16_t GetLevel() const;
		uint8_t GetIsFirstMessageInBatch() const;
		uint8_t GetIsLastMessageInBatch() const;
		t_DateTimeWithMilliseconds GetDateTime() const;
		uint32_t GetNumOrders() const;
	};

#pragma pack(push, 1)
	/*==========================================================================*/
	struct s_MarketDepthSnapshotLevelFloat
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DEPTH_SNAPSHOT_LEVEL_FLOAT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		float Price = 0;
		float Quantity = 0;
		uint32_t NumOrders = 0;
		uint16_t Level = 0;
		AtBidOrAskEnum8 Side = BID_ASK_UNSET_8;
		FinalUpdateInBatchEnum FinalUpdateInBatch = FINAL_UPDATE_UNSET;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetSymbolID() const;
		float GetPrice() const;
		float GetQuantity() const;
		uint32_t GetNumOrders() const;
		uint16_t GetLevel() const;
		AtBidOrAskEnum8 GetSide() const;
		FinalUpdateInBatchEnum GetFinalUpdateInBatch() const;
	};
#pragma pack(pop)
	
	/*==========================================================================*/
	struct s_MarketDepthUpdateLevel
	{	 
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DEPTH_UPDATE_LEVEL;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		AtBidOrAskEnum Side = BID_ASK_UNSET;
		double Price = 0;
		double Quantity = 0;
		MarketDepthUpdateTypeEnum UpdateType = MARKET_DEPTH_UNSET;
		t_DateTimeWithMilliseconds DateTime = 0;
		uint32_t NumOrders = 0;

		FinalUpdateInBatchEnum FinalUpdateInBatch = FINAL_UPDATE_UNSET;
		uint16_t Level = 0;
				
		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		AtBidOrAskEnum GetSide() const;
		double GetPrice() const;
		double GetQuantity() const;
		MarketDepthUpdateTypeEnum GetUpdateType() const;
		t_DateTimeWithMilliseconds GetDateTime() const;
		uint32_t GetNumOrders() const;
		FinalUpdateInBatchEnum GetFinalUpdateInBatch() const;
		uint16_t GetLevel() const;
	};

	/*==========================================================================*/
	#pragma pack(push, 1)
	struct s_MarketDepthUpdateLevelFloatWithMilliseconds
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DEPTH_UPDATE_LEVEL_FLOAT_WITH_MILLISECONDS;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		t_DateTimeWithMillisecondsInt DateTime = 0;
		float Price = 0;
		float Quantity = 0;
		AtBidOrAskEnum8 Side = BID_ASK_UNSET_8;
		MarketDepthUpdateTypeEnum UpdateType = MARKET_DEPTH_UNSET;
		uint16_t NumOrders = 0;
		FinalUpdateInBatchEnum FinalUpdateInBatch = FINAL_UPDATE_UNSET;
		uint16_t Level = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetSymbolID() const;
		t_DateTimeWithMillisecondsInt GetDateTime() const;
		float GetPrice() const;
		float GetQuantity() const;
		int8_t GetSide() const;
		int8_t GetUpdateType() const;
		uint16_t GetNumOrders() const;
		FinalUpdateInBatchEnum GetFinalUpdateInBatch() const;
		uint16_t GetLevel() const;
	};

	/*==========================================================================*/
	struct s_MarketDepthUpdateLevelNoTimestamp
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DEPTH_UPDATE_LEVEL_NO_TIMESTAMP;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		float Price = 0;
		float Quantity = 0;
		uint16_t NumOrders = 0;
		int8_t Side = 0;
		int8_t UpdateType = 0;
		FinalUpdateInBatchEnum FinalUpdateInBatch = FINAL_UPDATE_UNSET;
		uint16_t Level = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetSymbolID() const;
		float GetPrice() const;
		float GetQuantity() const;
		uint16_t GetNumOrders() const;
		int8_t GetSide() const;
		int8_t GetUpdateType() const;
		FinalUpdateInBatchEnum GetFinalUpdateInBatch() const;
		uint16_t GetLevel() const;
	};

	/*==========================================================================*/
	struct s_MarketDataUpdateTradeNoTimestamp
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_TRADE_NO_TIMESTAMP;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		float Price = 0;
		uint32_t Volume = 0;
		AtBidOrAskEnum8 AtBidOrAsk = BID_ASK_UNSET_8;
		UnbundledTradeIndicatorEnum UnbundledTradeIndicator = UNBUNDLED_TRADE_NONE;
		TradeConditionEnum TradeCondition = TRADE_CONDITION_NONE;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		float GetPrice() const;
		uint32_t GetVolume() const;
		AtBidOrAskEnum8 GetAtBidOrAsk() const;
		UnbundledTradeIndicatorEnum GetUnbundledTradeIndicator() const;
		TradeConditionEnum GetTradeCondition() const;
	};


	#pragma pack(pop)

	/*==========================================================================*/
	struct s_MarketDataUpdateSessionSettlement
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_SESSION_SETTLEMENT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		double Price = 0;
		t_DateTime4Byte DateTime = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		double GetPrice() const;
		t_DateTime4Byte GetDateTime() const;
	};

	/*==========================================================================*/
	struct s_MarketDataUpdateSessionOpen
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_SESSION_OPEN;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		double Price = 0;
		t_DateTime4Byte TradingSessionDate = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		double GetPrice() const;
		t_DateTime4Byte GetTradingSessionDate() const;
	};

	/*==========================================================================*/
	struct s_MarketDataUpdateSessionNumTrades
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_SESSION_NUM_TRADES;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		int32_t NumTrades = 0;

		t_DateTime4Byte TradingSessionDate = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetSymbolID() const;
		int32_t GetNumTrades() const;
		t_DateTime4Byte GetTradingSessionDate() const;
	};

	/*==========================================================================*/
	struct s_MarketDataUpdateTradingSessionDate
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_TRADING_SESSION_DATE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		t_DateTime4Byte Date = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetSymbolID() const;
		t_DateTime4Byte GetDate() const;
	};

	/*==========================================================================*/
	struct s_MarketDepthReject
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DEPTH_REJECT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;
		uint32_t SymbolID = 0;
		char RejectText[TEXT_DESCRIPTION_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		const char* GetRejectText();
		void SetRejectText(const char* NewValue);
	};

	/*==========================================================================*/
	struct s_MarketDataUpdateTrade
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_TRADE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;

		AtBidOrAskEnum AtBidOrAsk = BID_ASK_UNSET;

		double Price = 0;
		double Volume = 0;
		t_DateTimeWithMilliseconds DateTime = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		AtBidOrAskEnum GetAtBidOrAsk() const;
		double GetPrice() const;
		double GetVolume() const;
		t_DateTimeWithMilliseconds GetDateTime() const;

	};

	/*==========================================================================*/
	struct s_MarketDataUpdateTradeWithUnbundledIndicator
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		AtBidOrAskEnum8 AtBidOrAsk = BID_ASK_UNSET_8;
		UnbundledTradeIndicatorEnum UnbundledTradeIndicator = UNBUNDLED_TRADE_NONE;
		uint8_t TradeCondition = 0;
		uint8_t Reserve_1 = 0;
		uint32_t Reserve_2 = 0;
		double Price = 0;
		uint32_t Volume = 0;
		uint32_t Reserve_3 = 0;
		t_DateTimeWithMilliseconds DateTime = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		AtBidOrAskEnum8 GetAtBidOrAsk() const;
		UnbundledTradeIndicatorEnum GetUnbundledTradeIndicator() const;
		uint8_t GetTradeCondition() const;
		double GetPrice() const;
		uint32_t GetVolume() const;
		t_DateTimeWithMilliseconds GetDateTime() const;
	};

#pragma pack(push, 1)
	/*==========================================================================*/
	struct s_MarketDataUpdateTradeWithUnbundledIndicator2
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_2;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		float Price = 0;
		uint32_t Volume = 0;

		t_DateTimeWithMicrosecondsInt DateTime = 0;

		AtBidOrAskEnum8 AtBidOrAsk = BID_ASK_UNSET_8;
		UnbundledTradeIndicatorEnum UnbundledTradeIndicator = UNBUNDLED_TRADE_NONE;
		TradeConditionEnum TradeCondition = TRADE_CONDITION_NONE;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetSymbolID() const;
		float GetPrice() const;
		uint32_t GetVolume() const;
		t_DateTimeWithMicrosecondsInt GetDateTime() const;

		AtBidOrAskEnum8 GetAtBidOrAsk() const;
		UnbundledTradeIndicatorEnum GetUnbundledTradeIndicator() const;
		TradeConditionEnum GetTradeCondition() const;
	};

#pragma pack(pop)

	/*==========================================================================*/
	struct s_MarketDataUpdateBidAsk
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_BID_ASK;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;

		double BidPrice = DBL_MAX;
		float BidQuantity = 0;
		double AskPrice = DBL_MAX;
		float AskQuantity = 0;
		t_DateTime4Byte DateTime = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		double GetBidPrice() const;
		float GetBidQuantity() const;
		double GetAskPrice() const;
		float GetAskQuantity() const;
		t_DateTime4Byte GetDateTime() const;
	};

	/*==========================================================================*/
	struct s_MarketDataUpdateBidAskCompact
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_BID_ASK_COMPACT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		float BidPrice = FLT_MAX;
		float BidQuantity = 0;
		float AskPrice = FLT_MAX;
		float AskQuantity = 0;

		t_DateTime4Byte DateTime = 0;

		uint32_t SymbolID = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		float GetBidPrice() const;
		float GetBidQuantity() const;
		float GetAskPrice() const;
		float GetAskQuantity() const;
		t_DateTime4Byte GetDateTime() const;
		uint32_t GetSymbolID() const;
	};

#pragma pack(push, 1)
	/*==========================================================================*/
	struct s_MarketDataUpdateBidAskFloatWithMicroseconds
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_BID_ASK_FLOAT_WITH_MICROSECONDS;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		float BidPrice = FLT_MAX;
		float BidQuantity = 0;
		float AskPrice = FLT_MAX;
		float AskQuantity = 0;
		t_DateTimeWithMicrosecondsInt DateTime = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		float GetBidPrice() const;
		float GetBidQuantity() const;
		float GetAskPrice() const;
		float GetAskQuantity() const;
		t_DateTimeWithMicrosecondsInt GetDateTime() const;
	};

	/*==========================================================================*/
	struct s_MarketDataUpdateBidAskNoTimeStamp
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_BID_ASK_NO_TIMESTAMP;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;

		float BidPrice = FLT_MAX;
		uint32_t BidQuantity = 0;
		float AskPrice = FLT_MAX;
		uint32_t AskQuantity = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint16_t GetSize() const;
		uint16_t GetType() const;

		uint32_t GetSymbolID() const;

		float GetBidPrice() const;
		uint32_t GetBidQuantity() const;
		float GetAskPrice() const;
		uint32_t GetAskQuantity() const;
	};

#pragma pack(pop)

	/*==========================================================================*/
	struct s_MarketDataUpdateTradeCompact
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_TRADE_COMPACT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;
	
		float Price = 0;
		float Volume = 0;
		t_DateTime4Byte DateTime = 0;
		uint32_t SymbolID = 0;
		AtBidOrAskEnum AtBidOrAsk = BID_ASK_UNSET;
				
		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		float GetPrice() const;
		float GetVolume() const;
		t_DateTime4Byte GetDateTime() const;
		uint32_t GetSymbolID() const;
		AtBidOrAskEnum GetAtBidOrAsk() const;
	};

	/*==========================================================================*/
	struct s_MarketDataUpdateSessionVolume
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_SESSION_VOLUME;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		double Volume = 0;
		t_DateTime4Byte TradingSessionDate = 0;
		uint8_t IsFinalSessionVolume = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		double GetVolume() const;
		t_DateTime4Byte GetTradingSessionDate() const;
		uint8_t GetIsFinalSessionVolume() const;
	};
	/*==========================================================================*/
	struct s_MarketDataUpdateOpenInterest
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_OPEN_INTEREST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		uint32_t OpenInterest = 0;

		t_DateTime4Byte TradingSessionDate = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		uint32_t GetOpenInterest() const;
		t_DateTime4Byte GetTradingSessionDate() const;
	};
	/*==========================================================================*/
	struct s_MarketDataUpdateSessionHigh
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_SESSION_HIGH;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		double Price = 0;
		t_DateTime4Byte TradingSessionDate = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		double GetPrice() const;
		t_DateTime4Byte GetTradingSessionDate() const;
	};

	/*==========================================================================*/
	struct s_MarketDataUpdateSessionLow
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_SESSION_LOW;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		double Price = 0;
		t_DateTime4Byte TradingSessionDate = 0;
				
		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetSymbolID() const;
		double GetPrice() const;
		t_DateTime4Byte GetTradingSessionDate() const;
	};

	/*==========================================================================*/
	struct s_MarketDataUpdateLastTradeSnapshot
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_DATA_UPDATE_LAST_TRADE_SNAPSHOT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		double LastTradePrice = 0;
		double LastTradeVolume = 0;
		t_DateTimeWithMilliseconds LastTradeDateTime = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetSymbolID() const;
		double GetLastTradePrice() const;
		double GetLastTradeVolume() const;
		t_DateTimeWithMilliseconds GetLastTradeDateTime() const;
	};

#pragma pack(push, 1)
	/*==========================================================================*/
	struct s_MarketOrdersRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_ORDERS_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		RequestActionEnum RequestAction = SUBSCRIBE;
		uint32_t SymbolID = 0;
		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};
		uint32_t SendQuantitiesGreaterOrEqualTo = 0;


		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		RequestActionEnum GetRequestAction() const;
		uint32_t GetSymbolID() const;
		const char* GetSymbol();
		void SetSymbol(const char* NewValue);
		const char* GetExchange();
		void SetExchange(const char* NewValue);
		int32_t GetSendQuantitiesGreaterOrEqualTo() const;
	};

	/*==========================================================================*/
	struct s_MarketOrdersReject
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_ORDERS_REJECT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		char RejectText[TEXT_DESCRIPTION_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetSymbolID() const;
		const char* GetRejectText();
		void SetRejectText(const char* NewValue);
	};

	/*==========================================================================*/
	struct s_MarketOrdersAdd//MARKET_ORDERS_ADD
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_ORDERS_ADD;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		DTC::BuySellEnum Side = BUY_SELL_UNSET;
		uint32_t Quantity = 0;
		double Price = 0;
		uint64_t OrderID = 0;
		t_DateTimeWithMicrosecondsInt DateTime = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);
		
		uint32_t GetSymbolID() const;
		DTC::BuySellEnum GetSide() const;
		uint32_t GetQuantity() const;
		double GetPrice() const;
		uint64_t GetOrderID() const;
		t_DateTimeWithMicrosecondsInt GetDateTime() const;
	};

	/*==========================================================================*/
	struct s_MarketOrdersModify//MARKET_ORDERS_MODIFY
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_ORDERS_MODIFY;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		DTC::BuySellEnum Side = DTC::BUY_SELL_UNSET;
		uint32_t Quantity = 0;
		double Price = 0;
		uint64_t OrderID = 0;
		double PriorPrice = 0;
		uint32_t PriorQuantity = 0;
		uint64_t PriorOrderID = 0;
		t_DateTimeWithMicrosecondsInt DateTime = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetSymbolID() const;
		DTC::BuySellEnum GetSide() const;
		uint32_t GetQuantity() const;
		double GetPrice() const;
		uint64_t GetOrderID() const;
		double GetPriorPrice() const;
		uint32_t GetPriorQuantity() const;
		uint64_t GetPriorOrderID() const;
		t_DateTimeWithMicrosecondsInt GetDateTime() const;

	};

	/*==========================================================================*/
	struct s_MarketOrdersRemove//MARKET_ORDERS_REMOVE
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_ORDERS_REMOVE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		uint32_t Quantity = 0;
		double Price = 0;
		uint64_t OrderID = 0;
		t_DateTimeWithMicrosecondsInt DateTime = 0;
		DTC::BuySellEnum Side = BUY_SELL_UNSET;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetSymbolID() const;
		uint32_t GetQuantity() const;
		double GetPrice() const;
		uint64_t GetOrderID() const;
		t_DateTimeWithMicrosecondsInt GetDateTime() const;
		DTC::BuySellEnum GetSide() const;

	};

	/*==========================================================================*/
	
	struct s_MarketOrdersSnapshotMessageBoundary
	{
		static constexpr uint16_t MESSAGE_TYPE = MARKET_ORDERS_SNAPSHOT_MESSAGE_BOUNDARY;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		uint32_t SymbolID = 0;
		MessageSetBoundaryEnum MessageBoundary = MESSAGE_SET_BOUNDARY_UNSET;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetSymbolID() const;
		MessageSetBoundaryEnum GetMessageBoundary() const;

	};

#pragma pack(pop)
	/*==========================================================================*/
	struct s_SubmitNewSingleOrder
	{
		static constexpr uint16_t MESSAGE_TYPE = SUBMIT_NEW_SINGLE_ORDER;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};

		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};
		char ClientOrderID[ORDER_ID_LENGTH] = {};

		OrderTypeEnum OrderType = ORDER_TYPE_UNSET;

		BuySellEnum BuySell = BUY_SELL_UNSET;

		double Price1 = 0;
		double Price2 = 0;
		double Quantity = 0;

		TimeInForceEnum TimeInForce = TIF_UNSET;

		t_DateTime GoodTillDateTime = 0;		
		uint8_t IsAutomatedOrder = 0;
		uint8_t IsParentOrder = 0;

		char FreeFormText[ORDER_FREE_FORM_TEXT_LENGTH] = {};

		OpenCloseTradeEnum OpenOrClose = TRADE_UNSET;
		double MaxShowQuantity = 0;

		char Price1AsString[PRICE_STRING_LENGTH] = {};
		char Price2AsString[PRICE_STRING_LENGTH] = {};

		double IntendedPositionQuantity = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		const char* GetSymbol();
		void SetSymbol(const char* NewValue);
		const char* GetExchange();
		void SetExchange(const char* NewValue);
		const char* GetTradeAccount();
		void SetTradeAccount(const char* NewValue);
		const char* GetClientOrderID();
		void SetClientOrderID(const char* NewValue);
		OrderTypeEnum GetOrderType() const;
		BuySellEnum GetBuySell() const;
		double GetPrice1() const;
		double GetPrice2() const;
		double GetQuantity() const;
		TimeInForceEnum GetTimeInForce() const;
		t_DateTime GetGoodTillDateTime() const;
		uint8_t GetIsAutomatedOrder() const;
		uint8_t GetIsParentOrder() const;
		const char* GetFreeFormText();
		void SetFreeFormText(const char* NewValue);
		OpenCloseTradeEnum GetOpenOrClose() const;
		double GetMaxShowQuantity() const;

		const char* GetPrice1AsString();
		void SetPrice1AsString(const char* NewValue);
		const char* GetPrice2AsString();
		void SetPrice2AsString(const char* NewValue);

		double GetIntendedPositionQuantity() const;
	};

	/*==========================================================================*/
	struct s_SubmitFlattenPositionOrder
	{
		static constexpr uint16_t MESSAGE_TYPE = SUBMIT_FLATTEN_POSITION_ORDER;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};

		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};
		char ClientOrderID[ORDER_ID_LENGTH] = {};
		char FreeFormText[ORDER_FREE_FORM_TEXT_LENGTH] = {};
		uint8_t IsAutomatedOrder = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		const char* GetSymbol();
		void SetSymbol(const char* NewValue);

		const char* GetExchange();
		void SetExchange(const char* NewValue);

		const char* GetTradeAccount();
		void SetTradeAccount(const char* NewValue);

		const char* GetClientOrderID();
		void SetClientOrderID(const char* NewValue);

		const char* GetFreeFormText();
		void SetFreeFormText(const char* NewValue);

		uint8_t GetIsAutomatedOrder() const;
	};

	/*==========================================================================*/
	struct s_CancelReplaceOrder
	{
		static constexpr uint16_t MESSAGE_TYPE = CANCEL_REPLACE_ORDER;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		char ServerOrderID[ORDER_ID_LENGTH] = {};
		char ClientOrderID[ORDER_ID_LENGTH] = {};

		double Price1 = 0;
		double Price2 = 0;
		double Quantity = 0;
		uint8_t Price1IsSet = 1;
		uint8_t Price2IsSet = 1;

		int32_t Unused = 0;
		TimeInForceEnum TimeInForce = TIF_UNSET;
		t_DateTime GoodTillDateTime = 0;
		uint8_t UpdatePrice1OffsetToParent = 0;
		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};

		char Price1AsString[PRICE_STRING_LENGTH] = {};
		char Price2AsString[PRICE_STRING_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		const char* GetServerOrderID();
		void SetServerOrderID(const char* NewValue);
		const char* GetClientOrderID();
		void SetClientOrderID(const char* NewValue);
		const char* GetTradeAccount();
		void SetTradeAccount(const char* NewValue);
		double GetPrice1() const;
		double GetPrice2() const;
		double GetQuantity() const;
		uint8_t GetPrice1IsSet() const;
		uint8_t GetPrice2IsSet() const;
		TimeInForceEnum GetTimeInForce() const;
		t_DateTime GetGoodTillDateTime() const;
		uint8_t GetUpdatePrice1OffsetToParent() const;

		const char* GetPrice1AsString();
		void SetPrice1AsString(const char* NewValue);

		const char* GetPrice2AsString();
		void SetPrice2AsString(const char* NewValue);
	};

	/*==========================================================================*/
	struct s_CancelOrder
	{
		static constexpr uint16_t MESSAGE_TYPE = CANCEL_ORDER;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;
		
		char ServerOrderID[ORDER_ID_LENGTH] = {};
		char ClientOrderID[ORDER_ID_LENGTH] = {};
		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		const char* GetServerOrderID();
		void SetServerOrderID(const char* NewValue);
		const char* GetClientOrderID();
		void SetClientOrderID(const char* NewValue);
		const char* GetTradeAccount();
		void SetTradeAccount(const char* NewValue);
	};

	/*==========================================================================*/
	struct s_SubmitNewOCOOrder
	{
		static constexpr uint16_t MESSAGE_TYPE = SUBMIT_NEW_OCO_ORDER;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};

		char ClientOrderID_1[ORDER_ID_LENGTH] = {};
		OrderTypeEnum OrderType_1 = ORDER_TYPE_UNSET;
		BuySellEnum BuySell_1 = BUY_SELL_UNSET;
		double Price1_1 = 0;
		double Price2_1 = 0;
		double Quantity_1 = 0;

		char ClientOrderID_2[ORDER_ID_LENGTH] = {};
		OrderTypeEnum OrderType_2 = ORDER_TYPE_UNSET;
		BuySellEnum BuySell_2 = BUY_SELL_UNSET;
		double Price1_2 = 0;
		double Price2_2 = 0;
		double Quantity_2 = 0;

		TimeInForceEnum TimeInForce = TIF_UNSET;
		t_DateTime GoodTillDateTime = 0;

		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};

		uint8_t IsAutomatedOrder = 0;

		char ParentTriggerClientOrderID[ORDER_ID_LENGTH] = {};

		char FreeFormText[ORDER_FREE_FORM_TEXT_LENGTH] = {};

		OpenCloseTradeEnum OpenOrClose = TRADE_UNSET;

		PartialFillHandlingEnum PartialFillHandling = PARTIAL_FILL_UNSET;

		uint8_t UseOffsets = 0;
		double OffsetFromParent1 = 0;
		double OffsetFromParent2 = 0;
		uint8_t MaintainSamePricesOnParentFill = 0;

		char Price1_1AsString[PRICE_STRING_LENGTH] = {};
		char Price2_1AsString[PRICE_STRING_LENGTH] = {};
		char Price1_2AsString[PRICE_STRING_LENGTH] = {};
		char Price2_2AsString[PRICE_STRING_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		void SetClientOrderID_1(const char* NewValue);
		void SetClientOrderID_2(const char* NewValue);
		const char* GetFreeFormText();
		void SetFreeFormText(const char* NewValue);
		const char* GetClientOrderID_1();
		const char* GetClientOrderID_2();
		const char* GetSymbol();
		void SetSymbol(const char* NewValue);
		const char* GetExchange();
		void SetExchange(const char* NewValue);
		OrderTypeEnum GetOrderType_1() const;
		OrderTypeEnum GetOrderType_2() const;
		BuySellEnum GetBuySell_1() const;
		BuySellEnum GetBuySell_2() const;
		TimeInForceEnum GetTimeInForce() const;
		t_DateTime GetGoodTillDateTime() const;
		void SetParentTriggerClientOrderID(const char* NewValue);
		const char* GetParentTriggerClientOrderID();
		uint8_t GetIsAutomatedOrder() const;
		double GetPrice1_1() const;
		double GetPrice2_1() const;
		double GetPrice1_2() const;
		double GetPrice2_2() const;
		double GetQuantity_1() const;
		double GetQuantity_2() const;
		const char* GetTradeAccount();
		void SetTradeAccount(const char* NewValue);
		OpenCloseTradeEnum GetOpenOrClose() const;
		PartialFillHandlingEnum GetPartialFillHandling() const;
		uint8_t GetUseOffsets() const;
		double GetOffsetFromParent1() const;
		double GetOffsetFromParent2() const;
		uint8_t GetMaintainSamePricesOnParentFill() const;

		const char* GetPrice1_1AsString();
		void SetPrice1_1AsString(const char* NewValue);
		const char* GetPrice2_1AsString();
		void SetPrice2_1AsString(const char* NewValue);
		const char* GetPrice1_2AsString();
		void SetPrice1_2AsString(const char* NewValue);
		const char* GetPrice2_2AsString();
		void SetPrice2_2AsString(const char* NewValue);
	};

	/*==========================================================================*/
	struct s_OpenOrdersRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = OPEN_ORDERS_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		int32_t RequestAllOrders = 1;

		char ServerOrderID[ORDER_ID_LENGTH] = {};
		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		int32_t GetRequestAllOrders() const;
		void SetServerOrderID(const char* NewValue);
		const char* GetServerOrderID();
		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();
	};

	/*==========================================================================*/
	struct s_HistoricalOrderFillsRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_ORDER_FILLS_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		char ServerOrderID[ORDER_ID_LENGTH] = {};

		int32_t NumberOfDays = 0;

		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};

		t_DateTime StartDateTime = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		int32_t GetNumberOfDays() const;
		t_DateTime GetStartDateTime() const;
		void SetServerOrderID(const char* NewValue);
		const char* GetServerOrderID();
		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();
	};

	/*==========================================================================*/
	struct s_HistoricalOrderFillsReject
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_ORDER_FILLS_REJECT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char RejectText[TEXT_DESCRIPTION_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		void SetRejectText(const char* NewValue);
		const char* GetRejectText();
	};


	/*==========================================================================*/
	struct s_CurrentPositionsRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = CURRENT_POSITIONS_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char  TradeAccount[TRADE_ACCOUNT_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();
	};

	/*==========================================================================*/
	struct s_CurrentPositionsReject
	{
		static constexpr uint16_t MESSAGE_TYPE = CURRENT_POSITIONS_REJECT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char RejectText[TEXT_DESCRIPTION_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		void SetRejectText(const char* NewValue);
		const char* GetRejectText();
	};

	/*==========================================================================*/
	struct s_OrderUpdate
	{
		static constexpr uint16_t MESSAGE_TYPE = ORDER_UPDATE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		int32_t TotalNumMessages = 0;
		int32_t MessageNumber = 0;

		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};

		char PreviousServerOrderID[ORDER_ID_LENGTH] = {};

		char ServerOrderID[ORDER_ID_LENGTH] = {};

		char ClientOrderID[ORDER_ID_LENGTH] = {};

		char ExchangeOrderID[ORDER_ID_LENGTH] = {};

		OrderStatusEnum OrderStatus = ORDER_STATUS_UNSPECIFIED;

		OrderUpdateReasonEnum OrderUpdateReason = ORDER_UPDATE_REASON_UNSET;

		OrderTypeEnum OrderType = ORDER_TYPE_UNSET;

		BuySellEnum BuySell = BUY_SELL_UNSET;

		double Price1 = DBL_MAX;
		double Price2 = DBL_MAX;

		TimeInForceEnum TimeInForce = TIF_UNSET;
		t_DateTime GoodTillDateTime = 0;
		double OrderQuantity = DBL_MAX;
		double FilledQuantity = DBL_MAX;
		double RemainingQuantity = DBL_MAX;
		double AverageFillPrice = DBL_MAX;

		double LastFillPrice = DBL_MAX;
		t_DateTimeWithMillisecondsInt LastFillDateTime = 0;
		double LastFillQuantity = DBL_MAX;
		char LastFillExecutionID[64] = {};

		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};
		char InfoText[TEXT_DESCRIPTION_LENGTH] = {};

		uint8_t NoOrders = 0;
		char ParentServerOrderID[ORDER_ID_LENGTH] = {};
		char OCOLinkedOrderServerOrderID[ORDER_ID_LENGTH] = {};

		OpenCloseTradeEnum OpenOrClose = TRADE_UNSET;

		char PreviousClientOrderID[ORDER_ID_LENGTH] = {};
		char FreeFormText[ORDER_FREE_FORM_TEXT_LENGTH] = {};
		t_DateTime OrderReceivedDateTime = 0;
		t_DateTimeWithMilliseconds LatestTransactionDateTime = 0;

		char Username[USERNAME_PASSWORD_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		const char* GetSymbol();
		void SetSymbol(const char* NewValue);
		const char* GetExchange();
		void SetExchange(const char* NewValue);
		const char* GetPreviousServerOrderID();
		void SetPreviousServerOrderID(const char* NewValue);
		const char* GetServerOrderID();
		void SetServerOrderID(const char* NewValue);
		const char* GetClientOrderID();
		void SetClientOrderID(const char* NewValue);
		const char* GetExchangeOrderID();
		void SetExchangeOrderID(const char* NewValue);
		void SetLastFillExecutionID(const char* NewValue);
		double GetOrderQuantity() const;
		double GetFilledQuantity() const;
		double GetRemainingQuantity() const;
		double GetLastFillQuantity() const;

		int32_t GetRequestID() const;
		int32_t GetMessageNumber() const;
		int32_t GetTotalNumMessages() const;
		OrderStatusEnum GetOrderStatus() const;
		OrderUpdateReasonEnum GetOrderUpdateReason() const;
		OrderTypeEnum GetOrderType() const;
		BuySellEnum GetBuySell() const;
		double GetPrice1() const;
		double GetPrice2() const;
		TimeInForceEnum GetTimeInForce() const;
		t_DateTime GetGoodTillDateTime() const;
		double GetAverageFillPrice() const;
		double GetLastFillPrice() const;
		t_DateTime GetLastFillDateTime() const;
		const char* GetLastFillExecutionID();

		const char* GetTradeAccount();
		void SetTradeAccount(const char* NewValue);
		const char* GetInfoText();
		void SetInfoText(const char* NewValue);
		uint8_t GetNoOrders() const;
		const char* GetParentServerOrderID();
		void SetParentServerOrderID(const char* NewValue);
		const char* GetOCOLinkedOrderServerOrderID();
		void SetOCOLinkedOrderServerOrderID(const char* NewValue);

		OpenCloseTradeEnum GetOpenOrClose() const;

		t_DateTime GetOrderReceivedDateTime() const;

		t_DateTimeWithMilliseconds GetLatestTransactionDateTime() const;

		const char* GetPreviousClientOrderID();
		void SetPreviousClientOrderID(const char* NewValue);

		const char* GetFreeFormText();
		void SetFreeFormText(const char* NewValue);

		const char* GetUsername();
		void SetUsername(const char* NewValue);
	};
	
	/*==========================================================================*/
	struct s_OpenOrdersReject
	{
		static constexpr uint16_t MESSAGE_TYPE = OPEN_ORDERS_REJECT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char RejectText[TEXT_DESCRIPTION_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		const char* GetRejectText();
		void SetRejectText(const char* NewValue);
	};

	/*==========================================================================*/
	struct s_HistoricalOrderFillResponse
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_ORDER_FILL_RESPONSE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0; 

		int32_t TotalNumberMessages = 0;
		int32_t MessageNumber = 0;

		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};
		char ServerOrderID[ORDER_ID_LENGTH] = {};
		BuySellEnum BuySell = BUY_SELL_UNSET;
		double Price = 0;
		t_DateTime DateTime = 0;
		double Quantity = 0;
		char UniqueExecutionID[ORDER_FILL_EXECUTION_LENGTH] = {};
		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};

		OpenCloseTradeEnum OpenClose = TRADE_UNSET;

		uint8_t NoOrderFills = 0;
		char InfoText[TEXT_DESCRIPTION_LENGTH] = {};

		double HighPriceDuringPosition = 0;
		double LowPriceDuringPosition = 0;
		double PositionQuantity = DBL_MAX;

		char Username[USERNAME_PASSWORD_LENGTH] = {};
		char ExchangeOrderID[ORDER_ID_LENGTH] = {};
		char SenderSubID[USERNAME_PASSWORD_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		int32_t GetMessageNumber() const;
		int32_t GetTotalNumberMessages() const;
		const char* GetSymbol();
		void SetSymbol(const char* NewValue);
		const char* GetExchange();
		void SetExchange(const char* NewValue);
		const char* GetServerOrderID();
		void SetServerOrderID(const char* NewValue);
		BuySellEnum GetBuySell() const;
		double GetPrice() const;
		t_DateTime GetDateTime() const;
		double GetQuantity() const;
		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();
		void SetUniqueExecutionID(const char* NewValue);
		const char* GetUniqueExecutionID();
		OpenCloseTradeEnum GetOpenClose() const;
		uint8_t GetNoOrderFills() const;
		double GetHighPriceDuringPosition() const;
		double GetLowPriceDuringPosition() const;
		double GetPositionQuantity() const;
		void SetInfoText(const char* NewValue);
		const char* GetInfoText();
		void SetUsername(const char* NewValue);
		const char* GetUsername();
		void SetExchangeOrderID(const char* NewValue);
		const char* GetExchangeOrderID();
		void SetSenderSubID(const char* NewValue);
		const char* GetSenderSubID();
	};

	/*==========================================================================*/
	struct s_PositionUpdate
	{
		static constexpr uint16_t MESSAGE_TYPE = POSITION_UPDATE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		int32_t TotalNumberMessages = 0;
		int32_t MessageNumber = 0;

		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};

		double Quantity = 0;
		double AveragePrice = 0;

		char PositionIdentifier[ORDER_ID_LENGTH] = {};

		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};
		uint8_t NoPositions = 0;
		uint8_t Unsolicited = 0;
		double MarginRequirement = 0;
		DTC::t_DateTime4Byte EntryDateTime = 0;
		double OpenProfitLoss = 0;
		double HighPriceDuringPosition = 0;
		double LowPriceDuringPosition = 0;
		double QuantityLimit = 0;
		double MaxPotentialPostionQuantity = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		int32_t GetTotalNumberMessages() const;
		int32_t GetMessageNumber() const;
		void SetSymbol(const char* NewValue);
		const char* GetSymbol();
		void SetExchange(const char* NewValue);
		const char* GetExchange();
		double GetQuantity() const;
		double GetAveragePrice() const;
		void SetPositionIdentifier(const char* NewValue);
		const char* GetPositionIdentifier();
		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();
		uint8_t GetNoPositions() const;
		uint8_t GetUnsolicited() const;
		double GetMarginRequirement() const;
		t_DateTime4Byte GetEntryDateTime() const;
		double GetOpenProfitLoss() const;
		double GetHighPriceDuringPosition() const;
		double GetLowPriceDuringPosition() const;

		double GetQuantityLimit() const;
		double GetMaxPotentialPostionQuantity() const;
	};

	/*==========================================================================*/
	struct s_AddCorrectingOrderFill
	{
		static constexpr uint16_t MESSAGE_TYPE = ADD_CORRECTING_ORDER_FILL;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};
		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};
		char ClientOrderID[ORDER_ID_LENGTH] = {};
		DTC::BuySellEnum BuySell = BUY_SELL_UNSET;
		double FillPrice = 0;
		double FillQuantity = 0;
		char FreeFormText[ORDER_FREE_FORM_TEXT_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		void SetSymbol(const char* NewValue);
		const char* GetSymbol();
		void SetExchange(const char* NewValue);
		const char* GetExchange();
		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();
		void SetClientOrderID(const char* NewValue);
		const char* GetClientOrderID();
		DTC::BuySellEnum GetBuySell() const;
		double GetFillPrice() const;
		double GetFillQuantity() const;
		void SetFreeFormText(const char* NewValue);
		const char* GetFreeFormText();
	};

	/*==========================================================================*/
	struct s_CorrectingOrderFillResponse
	{
		static constexpr uint16_t MESSAGE_TYPE = CORRECTING_ORDER_FILL_RESPONSE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		char ClientOrderID[ORDER_ID_LENGTH] = {};
		char ResultText[TEXT_MESSAGE_LENGTH] = {};
		uint8_t IsError = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		void SetClientOrderID(const char* NewValue);
		const char* GetClientOrderID();
		void SetResultText(const char* NewValue);
		const char* GetResultText();
		uint8_t GetIsError() const;
	};

	/*==========================================================================*/
	struct s_TradeAccountsRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = TRADE_ACCOUNTS_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
	};

	/*==========================================================================*/
	struct s_TradeAccountResponse
	{
		static constexpr uint16_t MESSAGE_TYPE = TRADE_ACCOUNT_RESPONSE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t TotalNumberMessages = 0;

		int32_t MessageNumber = 0;

		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};

		int32_t RequestID = 0;
				
		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetTotalNumberMessages() const;
		int32_t GetMessageNumber() const;
		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();
		int32_t GetRequestID() const;
	};

	/*==========================================================================*/
	struct s_ExchangeListRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = EXCHANGE_LIST_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
	};

	/*==========================================================================*/
	struct s_ExchangeListResponse
	{
		static constexpr uint16_t MESSAGE_TYPE = EXCHANGE_LIST_RESPONSE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char Exchange[EXCHANGE_LENGTH] = {};
		uint8_t IsFinalMessage = 0;
		char Description[EXCHANGE_DESCRIPTION_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		const char* GetExchange();
		void SetExchange(const char* NewValue);
		const char* GetDescription();
		void SetDescription(const char* NewValue);
		uint8_t GetIsFinalMessage() const;
	};

	/*==========================================================================*/
	struct s_SymbolsForExchangeRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = SYMBOLS_FOR_EXCHANGE_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char Exchange[EXCHANGE_LENGTH] = {};

		SecurityTypeEnum SecurityType = SECURITY_TYPE_UNSET;
		RequestActionEnum RequestAction = SUBSCRIBE;
		char Symbol[SYMBOL_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		const char* GetExchange();
		void SetExchange(const char* NewValue);
		SecurityTypeEnum GetSecurityType() const;
		RequestActionEnum GetRequestAction() const;
		const char* GetSymbol();
		void SetSymbol(const char* NewValue);
	};

	/*==========================================================================*/
	struct s_UnderlyingSymbolsForExchangeRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		char Exchange[EXCHANGE_LENGTH] = {};

		SecurityTypeEnum SecurityType = SECURITY_TYPE_UNSET;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		const char* GetExchange();
		void SetExchange(const char* NewValue);
		SecurityTypeEnum GetSecurityType() const;
	};

	/*==========================================================================*/
	struct s_SymbolsForUnderlyingRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = SYMBOLS_FOR_UNDERLYING_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		char UnderlyingSymbol[UNDERLYING_SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};

		SecurityTypeEnum SecurityType = SECURITY_TYPE_UNSET;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		const char* GetUnderlyingSymbol();
		void SetUnderlyingSymbol(const char* NewValue);
		const char* GetExchange();
		void SetExchange(const char* NewValue);
		SecurityTypeEnum GetSecurityType() const;
	};

	/*==========================================================================*/
	struct s_SymbolSearchRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = SYMBOL_SEARCH_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		char SearchText[SYMBOL_DESCRIPTION_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};

		SecurityTypeEnum SecurityType = SECURITY_TYPE_UNSET;
		SearchTypeEnum SearchType = SEARCH_TYPE_UNSET;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		const char* GetExchange();
		void SetExchange(const char* NewValue);
		const char* GetSearchText();
		void SetSearchText(const char* NewValue);
		SecurityTypeEnum GetSecurityType() const;
		SearchTypeEnum GetSearchType() const;
	};

	/*==========================================================================*/
	struct s_SecurityDefinitionForSymbolRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = SECURITY_DEFINITION_FOR_SYMBOL_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		void SetSymbol(const char* NewValue);
		const char* GetSymbol();
		void SetExchange(const char* NewValue);
		const char* GetExchange();
	};

	/*==========================================================================*/
	struct s_SecurityDefinitionResponse
	{
		static constexpr uint16_t MESSAGE_TYPE = SECURITY_DEFINITION_RESPONSE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};

		SecurityTypeEnum SecurityType = SECURITY_TYPE_UNSET;
		
		char Description[SYMBOL_DESCRIPTION_LENGTH] = {};
		float MinPriceIncrement = 0;
		PriceDisplayFormatEnum PriceDisplayFormat = PRICE_DISPLAY_FORMAT_UNSET;
		float CurrencyValuePerIncrement = 0;

		uint8_t IsFinalMessage = 0;

		float FloatToIntPriceMultiplier = 1.0;
		float IntToFloatPriceDivisor = 1.0;

		char UnderlyingSymbol[UNDERLYING_SYMBOL_LENGTH] = {};
		
		uint8_t UpdatesBidAskOnly = 0;

		float StrikePrice = 0;

		PutCallEnum PutOrCall = PC_UNSET;

		uint32_t ShortInterest = 0;

		t_DateTime4Byte SecurityExpirationDate = 0;

		float BuyRolloverInterest = 0;
		float SellRolloverInterest = 0;

		float EarningsPerShare = 0;
		uint32_t SharesOutstanding = 0;

		float IntToFloatQuantityDivisor = 0;

		uint8_t HasMarketDepthData = 1;

		float DisplayPriceMultiplier = 1.0;

		char ExchangeSymbol[SYMBOL_LENGTH] = {};

		float InitialMarginRequirement = 0;
		float MaintenanceMarginRequirement = 0;
		char Currency[CURRENCY_CODE_LENGTH] = {};

		float ContractSize = 0;
		uint32_t OpenInterest = 0;
		t_DateTime4Byte RolloverDate = 0;
		uint8_t IsDelayed = 0;
		int64_t SecurityIdentifier = 0;
		char ProductIdentifier[GENERAL_IDENTIFIER_LENGTH] = {};
				
		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		void SetSymbol(const char* NewValue);
		const char* GetSymbol();
		void SetExchange(const char* NewValue);
		const char* GetExchange();
		SecurityTypeEnum GetSecurityType() const;
		void SetDescription(const char* NewValue);
		const char* GetDescription();
		float GetMinPriceIncrement() const;
		PriceDisplayFormatEnum GetPriceDisplayFormat() const;
		float GetCurrencyValuePerIncrement() const;
		uint8_t GetIsFinalMessage() const;
		float GetFloatToIntPriceMultiplier() const;
		float GetIntToFloatPriceDivisor() const;
		const char* GetUnderlyingSymbol();
		void SetUnderlyingSymbol(const char* NewValue);

		uint8_t GetUpdatesBidAskOnly() const;
		float GetStrikePrice() const;
		PutCallEnum GetPutOrCall() const;
		uint32_t GetShortInterest() const;
		t_DateTime4Byte GetSecurityExpirationDate() const;
		float GetBuyRolloverInterest() const;
		float GetSellRolloverInterest() const;
		float GetEarningsPerShare() const;
		uint32_t GetSharesOutstanding() const;
		float GetIntToFloatQuantityDivisor() const;
		uint8_t GetHasMarketDepthData() const;
		float GetDisplayPriceMultiplier() const;
		const char* GetExchangeSymbol();
		void SetExchangeSymbol(const char* NewValue);
		float GetInitialMarginRequirement() const;
		float GetMaintenanceMarginRequirement() const;
		const char* GetCurrency();
		void SetCurrency(const char* NewValue);
		float GetContractSize() const;
		uint32_t GetOpenInterest() const;
		t_DateTime4Byte GetRolloverDate() const;
		uint8_t GetIsDelayed() const;
		int64_t GetSecurityIdentifier() const;

		const char* GetProductIdentifier();
		void SetProductIdentifier(const char* NewValue);
	};

	
	/*==========================================================================*/
	struct s_SecurityDefinitionReject
	{
		static constexpr uint16_t MESSAGE_TYPE = SECURITY_DEFINITION_REJECT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char RejectText[TEXT_DESCRIPTION_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetRequestID() const;

		void SetRejectText(const char* NewValue);
		const char* GetRejectText();
	};

	/*==========================================================================*/
	struct s_AccountBalanceRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = ACCOUNT_BALANCE_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;

		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();
	};

	/*==========================================================================*/
	struct s_AccountBalanceReject
	{
		static constexpr uint16_t MESSAGE_TYPE = ACCOUNT_BALANCE_REJECT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char RejectText[TEXT_DESCRIPTION_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetRequestID() const;

		void SetRejectText(const char* NewValue);
		const char* GetRejectText();
	};

	/*==========================================================================*/
	struct s_AccountBalanceUpdate
	{
		static constexpr uint16_t MESSAGE_TYPE = ACCOUNT_BALANCE_UPDATE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		double CashBalance = 0;

		double BalanceAvailableForNewPositions = 0;

		char AccountCurrency[8] = {};

		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};
		double SecuritiesValue = 0;
		double MarginRequirement = 0;

		int32_t TotalNumberMessages = 0;
		int32_t MessageNumber = 0;
		uint8_t NoAccountBalances = 0;
		uint8_t Unsolicited = 0;

		double OpenPositionsProfitLoss = 0;
		double DailyProfitLoss = 0;

		char InfoText[TEXT_DESCRIPTION_LENGTH] = {};
		uint64_t TransactionIdentifier = 0;
		double DailyNetLossLimit = 0;
		double TrailingAccountValueToLimitPositions = 0;

		uint8_t DailyNetLossLimitReached = 0;
		uint8_t IsUnderRequiredMargin = 0;
		uint8_t ClosePositionsAtEndOfDay = 0;
		uint8_t TradingIsDisabled = 0;
		char Description[TEXT_DESCRIPTION_LENGTH] = {};
		uint8_t IsUnderRequiredAccountValue = 0;
		t_DateTimeWithMicrosecondsInt TransactionDateTime = 0;

		double MarginRequirementFull = 0;
		double MarginRequirementFullPositionsOnly = 0;

		double PeakMarginRequirement = 0;
		char IntroducingBroker[TRADE_ACCOUNT_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		double GetCashBalance() const;
		double GetBalanceAvailableForNewPositions() const;
		void SetAccountCurrency(const char* NewValue);
		const char* GetAccountCurrency();
		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();
		double GetSecuritiesValue() const;
		double GetMarginRequirement() const;

		int32_t GetTotalNumberMessages() const;
		int32_t GetMessageNumber() const;
		uint8_t GetNoAccountBalances() const;
		uint8_t GetUnsolicited() const;

		double GetOpenPositionsProfitLoss() const;
		double GetDailyProfitLoss() const;

		void SetInfoText(const char* NewValue);
		const char* GetInfoText();

		uint64_t GetTransactionIdentifier() const;
		double GetDailyNetLossLimit() const;
		double GetTrailingAccountValueToLimitPositions() const;
		uint8_t GetDailyNetLossLimitReached() const;
		uint8_t GetIsUnderRequiredMargin() const;

		uint8_t GetClosePositionsAtEndOfDay() const;
		uint8_t GetTradingIsDisabled() const;

		void SetDescription(const char* NewValue);
		const char* GetDescription();

		uint8_t GetIsUnderRequiredAccountValue() const;
		t_DateTimeWithMicrosecondsInt GetTransactionDateTime() const;

		double GetMarginRequirementFull() const;
		double GetMarginRequirementFullPositionsOnly() const;
		double GetPeakMarginRequirement() const;

		void SetIntroducingBroker(const char* NewValue);
		const char* GetIntroducingBroker();

	};

	/*==========================================================================*/
	struct s_AccountBalanceAdjustment
	{
		static constexpr uint16_t MESSAGE_TYPE = ACCOUNT_BALANCE_ADJUSTMENT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};
		double CreditAmount = 0;
		double DebitAmount = 0;
		char Currency[CURRENCY_CODE_LENGTH] = {};

		char Reason[TEXT_DESCRIPTION_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;

		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();

		double GetCreditAmount() const;
		double GetDebitAmount() const;

		void SetCurrency(const char* NewValue);
		const char* GetCurrency();

		void SetReason(const char* NewValue);
		const char* GetReason();
	};

	/*==========================================================================*/
	struct s_AccountBalanceAdjustmentReject
	{
		static constexpr uint16_t MESSAGE_TYPE = ACCOUNT_BALANCE_ADJUSTMENT_REJECT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char RejectText[TEXT_DESCRIPTION_LENGTH] = {};
		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		void SetRejectText(const char* NewValue);
		const char* GetRejectText();

		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();
	};

	/*==========================================================================*/
	struct s_AccountBalanceAdjustmentComplete
	{
		static constexpr uint16_t MESSAGE_TYPE = ACCOUNT_BALANCE_ADJUSTMENT_COMPLETE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		int64_t TransactionID = 0;
		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};
		double NewBalance = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		int32_t GetRequestID() const;
		int64_t GetTransactionID() const;

		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();

		double GetNewBalance() const;
	};

	/*==========================================================================*/
	struct s_HistoricalAccountBalancesRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_ACCOUNT_BALANCES_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};

		t_DateTime StartDateTime = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;

		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();

		t_DateTime GetStartDateTime() const;
	};

	/*==========================================================================*/
	struct s_HistoricalAccountBalancesReject
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_ACCOUNT_BALANCES_REJECT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char RejectText[TEXT_DESCRIPTION_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetRequestID() const;

		void SetRejectText(const char* NewValue);
		const char* GetRejectText();
	};

	/*==========================================================================*/
	struct s_HistoricalAccountBalanceResponse
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_ACCOUNT_BALANCE_RESPONSE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		t_DateTimeWithMilliseconds DateTime = 0;
		double CashBalance = 0;
		char AccountCurrency[8] = {};
		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};
		uint8_t IsFinalResponse = 0;
		uint8_t NoAccountBalances = 0;
		char InfoText[TEXT_DESCRIPTION_LENGTH] = {};
		char TransactionId[GENERAL_IDENTIFIER_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		t_DateTimeWithMilliseconds GetDateTime() const;
		double GetCashBalance() const;
		void SetAccountCurrency(const char* NewValue);
		const char* GetAccountCurrency();
		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();
		uint8_t GetIsFinalResponse() const;
		uint8_t GetNoAccountBalances() const;
		void SetInfoText(const char* NewValue);
		const char* GetInfoText();
		void SetTransactionId(const char* NewValue);
		const char* GetTransactionId();
	};

	/*==========================================================================*/
	struct s_UserMessage
	{
		static constexpr uint16_t MESSAGE_TYPE = USER_MESSAGE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		char UserMessage[TEXT_MESSAGE_LENGTH] = {};

		uint8_t IsPopupMessage = 1;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		void SetUserMessage(const char* NewValue);
		const char* GetUserMessage();
		uint8_t GetIsPopupMessage() const;
	};

	/*==========================================================================*/
	struct s_GeneralLogMessage
	{
		static constexpr uint16_t MESSAGE_TYPE = GENERAL_LOG_MESSAGE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		char MessageText[128] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		void SetMessageText(const char* NewValue);
		const char* GetMessageText();
	};

	/*==========================================================================*/
	struct s_AlertMessage
	{
		static constexpr uint16_t MESSAGE_TYPE = ALERT_MESSAGE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		char MessageText[128] = {};
		char TradeAccount[TRADE_ACCOUNT_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		void SetMessageText(const char* NewValue);
		const char* GetMessageText();
		void SetTradeAccount(const char* NewValue);
		const char* GetTradeAccount();
	};

	/*==========================================================================*/
	struct s_JournalEntryAdd
	{
		static constexpr uint16_t MESSAGE_TYPE = JOURNAL_ENTRY_ADD;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		char JournalEntry[TEXT_MESSAGE_LENGTH] = {};
		t_DateTime DateTime = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		void SetJournalEntry(const char* NewValue);
		const char* GetJournalEntry();
		t_DateTime GetDateTime() const;
	};

	/*==========================================================================*/
	struct s_JournalEntriesRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = JOURNAL_ENTRIES_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		t_DateTime StartDateTime = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		t_DateTime GetStartDateTime() const;
	};

	/*==========================================================================*/
	struct s_JournalEntriesReject
	{
		static constexpr uint16_t MESSAGE_TYPE = JOURNAL_ENTRIES_REJECT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char RejectText[TEXT_DESCRIPTION_LENGTH] = {};

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		void SetRejectText(const char* NewValue);
		const char* GetRejectText();
	};

	/*==========================================================================*/
	struct s_JournalEntryResponse
	{
		static constexpr uint16_t MESSAGE_TYPE = JOURNAL_ENTRY_RESPONSE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		char JournalEntry[TEXT_MESSAGE_LENGTH] = {};
		t_DateTime DateTime = 0;
		uint8_t IsFinalResponse = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		void SetJournalEntry(const char* NewValue);
		const char* GetJournalEntry();
		t_DateTime GetDateTime() const;
		uint8_t GetIsFinalResponse() const;
	};
	
	/*==========================================================================*/
	struct s_HistoricalPriceDataRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_PRICE_DATA_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};
		HistoricalDataIntervalEnum RecordInterval = INTERVAL_TICK;
		t_DateTime StartDateTime = 0;
		t_DateTime EndDateTime = 0;
		uint32_t MaxDaysToReturn = 0;
		uint8_t  UseZLibCompression = 0;
		uint8_t RequestDividendAdjustedStockData = 0;
		uint16_t Integer_1 = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		void SetSymbol(const char* NewValue);
		const char* GetSymbol();
		void SetExchange(const char* NewValue);
		const char* GetExchange();
		HistoricalDataIntervalEnum GetRecordInterval() const;
		t_DateTime GetStartDateTime() const;
		t_DateTime GetEndDateTime() const;
		uint32_t GetMaxDaysToReturn() const;
		uint8_t GetUseZLibCompression() const;
		uint8_t GetRequestDividendAdjustedStockData() const;
		uint16_t GetInteger_1() const;
	};

	/*==========================================================================*/
	struct s_HistoricalPriceDataResponseHeader
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_PRICE_DATA_RESPONSE_HEADER;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		HistoricalDataIntervalEnum RecordInterval = INTERVAL_TICK;

		uint8_t UseZLibCompression = 0;
		
		uint8_t NoRecordsToReturn = 0;

		float IntToFloatPriceDivisor = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		HistoricalDataIntervalEnum GetRecordInterval() const;
		uint8_t GetUseZLibCompression() const;
		uint8_t GetNoRecordsToReturn() const;
		float GetIntToFloatPriceDivisor() const;
	};

	/*==========================================================================*/
	struct s_HistoricalPriceDataReject
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_PRICE_DATA_REJECT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char RejectText[TEXT_DESCRIPTION_LENGTH] = {};

		HistoricalPriceDataRejectReasonCodeEnum RejectReasonCode = HPDR_UNSET;
		uint16_t RetryTimeInSeconds = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		void SetRejectText(const char* NewValue);
		const char* GetRejectText();
		HistoricalPriceDataRejectReasonCodeEnum GetRejectReasonCode() const;
		uint16_t GetRetryTimeInSeconds() const;
	};

	/*==========================================================================*/
	struct s_HistoricalPriceDataRecordResponse
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_PRICE_DATA_RECORD_RESPONSE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		//Format can also be t_DateTime. Check value to determine. 
		t_DateTimeWithMicrosecondsInt StartDateTime = 0;

		double OpenPrice = 0;
		double HighPrice = 0;
		double LowPrice = 0;
		double LastPrice = 0;
		double Volume = 0;
		union
		{
			uint32_t OpenInterest = 0;
			uint32_t NumTrades;
		};
		double BidVolume = 0;
		double AskVolume = 0;

		uint8_t IsFinalRecord = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		t_DateTimeWithMicrosecondsInt GetStartDateTime() const;
		double GetOpenPrice() const;
		double GetHighPrice() const;
		double GetLowPrice() const;
		double GetLastPrice() const;
		double GetVolume() const;
		uint32_t GetOpenInterest() const;
		uint32_t GetNumTrades() const;
		double GetBidVolume() const;
		double GetAskVolume() const;
		uint8_t GetIsFinalRecord() const;
	};

	/*==========================================================================*/
	struct s_HistoricalPriceDataTickRecordResponse
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;
		int32_t RequestID = 0;

		t_DateTimeWithMilliseconds DateTime = 0;
		AtBidOrAskEnum AtBidOrAsk = BID_ASK_UNSET;

		double Price = 0;
		double Volume = 0;

		uint8_t IsFinalRecord = 0;
		
		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		t_DateTimeWithMilliseconds GetDateTime() const;
		double GetPrice() const;
		double GetVolume() const;
		AtBidOrAskEnum GetAtBidOrAsk() const;
		uint8_t GetIsFinalRecord() const;
	};

	/*==========================================================================*/
	struct s_HistoricalPriceDataResponseTrailer
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_PRICE_DATA_RESPONSE_TRAILER;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		t_DateTimeWithMicrosecondsInt FinalRecordLastDateTime = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		int32_t GetRequestID() const;
		t_DateTimeWithMicrosecondsInt GetFinalRecordLastDateTime() const;
	};

	/*==========================================================================*/
	struct s_HistoricalMarketDepthDataRequest
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_MARKET_DEPTH_DATA_REQUEST;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		char Symbol[SYMBOL_LENGTH] = {};
		char Exchange[EXCHANGE_LENGTH] = {};
		t_DateTimeWithMicrosecondsInt StartDateTime = 0;
		t_DateTimeWithMicrosecondsInt EndDateTime = 0;
		uint8_t  UseZLibCompression = 0;
		uint8_t Integer_1 = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		int32_t GetRequestID() const;
		void SetSymbol(const char* NewValue);
		const char* GetSymbol();
		void SetExchange(const char* NewValue);
		const char* GetExchange();
		t_DateTimeWithMicrosecondsInt GetStartDateTime() const;
		t_DateTimeWithMicrosecondsInt GetEndDateTime() const;
		uint8_t GetUseZLibCompression() const;
		uint8_t GetInteger_1() const;
	};

	/*==========================================================================*/
	struct s_HistoricalMarketDepthDataResponseHeader
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_MARKET_DEPTH_DATA_RESPONSE_HEADER;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;
		uint8_t UseZLibCompression = 0;
		uint8_t NoRecordsToReturn = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		int32_t GetRequestID() const;
		uint8_t GetUseZLibCompression() const;
		uint8_t GetNoRecordsToReturn() const;
	};

	/*==========================================================================*/
	struct s_HistoricalMarketDepthDataReject
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_MARKET_DEPTH_DATA_REJECT;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		char RejectText[TEXT_DESCRIPTION_LENGTH] = {};

		HistoricalPriceDataRejectReasonCodeEnum RejectReasonCode = DTC::HPDR_UNSET;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		int32_t GetRequestID() const;

		void SetRejectText(const char* NewValue);
		const char* GetRejectText();

		HistoricalPriceDataRejectReasonCodeEnum GetRejectReasonCode() const;
	};

	/*==========================================================================*/
	struct s_HistoricalMarketDepthDataRecordResponse
	{
		static constexpr uint16_t MESSAGE_TYPE = HISTORICAL_MARKET_DEPTH_DATA_RECORD_RESPONSE;

		uint16_t Size = sizeof(*this);
		uint16_t Type = MESSAGE_TYPE;

		int32_t RequestID = 0;

		t_DateTimeWithMicrosecondsInt StartDateTime = 0;
		uint8_t Command = 0;
		uint8_t Flags = 0;
		uint16_t NumOrders = 0;

		float Price = 0;
		uint32_t Quantity = 0;

		uint8_t IsFinalRecord = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);
		
		int32_t GetRequestID() const;

		t_DateTimeWithMicrosecondsInt GetStartDateTime() const;
		uint8_t GetCommand() const;
		uint8_t GetFlags() const;
		uint16_t GetNumOrders() const;

		float GetPrice() const;
		uint32_t GetQuantity() const;

		uint8_t GetIsFinalRecord() const;
	};

#pragma pack(pop)
}
