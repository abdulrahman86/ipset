// WARNING: This file has automatically been generated on Sun, 16 Dec 2018 22:49:13 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package ipset

const (
	// IPSET_PROTOCOL as defined in ipset/ipset.h:16
	IPSET_PROTOCOL = 7
	// IPSET_PROTOCOL_MIN as defined in ipset/ipset.h:17
	IPSET_PROTOCOL_MIN = 6
	// IPSET_MAXNAMELEN as defined in ipset/ipset.h:20
	IPSET_MAXNAMELEN = 32
	// IPSET_MAX_COMMENT_SIZE as defined in ipset/ipset.h:23
	IPSET_MAX_COMMENT_SIZE = 255
	// IPSET_ATTR_CMD_MAX as defined in ipset/ipset.h:74
	IPSET_ATTR_CMD_MAX = (__IPSET_ATTR_CMD_MAX - 1)
	// IPSET_ATTR_CREATE_MAX as defined in ipset/ipset.h:108
	IPSET_ATTR_CREATE_MAX = (__IPSET_ATTR_CREATE_MAX - 1)
	// IPSET_ATTR_ADT_MAX as defined in ipset/ipset.h:128
	IPSET_ATTR_ADT_MAX = (__IPSET_ATTR_ADT_MAX - 1)
	// IPSET_ATTR_IPADDR_MAX as defined in ipset/ipset.h:136
	IPSET_ATTR_IPADDR_MAX = (__IPSET_ATTR_IPADDR_MAX - 1)
	// IPSET_INVALID_ID as defined in ipset/ipset.h:233
	IPSET_INVALID_ID = 65535
)

// ipset_cmd as declared in ipset/ipset.h:26
type ipset_cmd int32

// ipset_cmd enumeration from ipset/ipset.h:26
const (
	IPSET_CMD_NONE        = iota
	IPSET_CMD_PROTOCOL    = 1
	IPSET_CMD_CREATE      = 2
	IPSET_CMD_DESTROY     = 3
	IPSET_CMD_FLUSH       = 4
	IPSET_CMD_RENAME      = 5
	IPSET_CMD_SWAP        = 6
	IPSET_CMD_LIST        = 7
	IPSET_CMD_SAVE        = 8
	IPSET_CMD_ADD         = 9
	IPSET_CMD_DEL         = 10
	IPSET_CMD_TEST        = 11
	IPSET_CMD_HEADER      = 12
	IPSET_CMD_TYPE        = 13
	IPSET_CMD_GET_BYNAME  = 14
	IPSET_CMD_GET_BYINDEX = 15
	IPSET_MSG_MAX         = 16
	IPSET_CMD_RESTORE     = IPSET_MSG_MAX
	IPSET_CMD_HELP        = 17
	IPSET_CMD_VERSION     = 18
	IPSET_CMD_QUIT        = 19
	IPSET_CMD_MAX         = 20
	IPSET_CMD_COMMIT      = IPSET_CMD_MAX
)

const (
	// IPSET_ATTR_UNSPEC as declared in ipset/ipset.h:58
	IPSET_ATTR_UNSPEC = iota
	// IPSET_ATTR_PROTOCOL as declared in ipset/ipset.h:59
	IPSET_ATTR_PROTOCOL = 1
	// IPSET_ATTR_SETNAME as declared in ipset/ipset.h:60
	IPSET_ATTR_SETNAME = 2
	// IPSET_ATTR_TYPENAME as declared in ipset/ipset.h:61
	IPSET_ATTR_TYPENAME = 3
	// IPSET_ATTR_SETNAME2 as declared in ipset/ipset.h:62
	IPSET_ATTR_SETNAME2 = IPSET_ATTR_TYPENAME
	// IPSET_ATTR_REVISION as declared in ipset/ipset.h:63
	IPSET_ATTR_REVISION = 4
	// IPSET_ATTR_FAMILY as declared in ipset/ipset.h:64
	IPSET_ATTR_FAMILY = 5
	// IPSET_ATTR_FLAGS as declared in ipset/ipset.h:65
	IPSET_ATTR_FLAGS = 6
	// IPSET_ATTR_DATA as declared in ipset/ipset.h:66
	IPSET_ATTR_DATA = 7
	// IPSET_ATTR_ADT as declared in ipset/ipset.h:67
	IPSET_ATTR_ADT = 8
	// IPSET_ATTR_LINENO as declared in ipset/ipset.h:68
	IPSET_ATTR_LINENO = 9
	// IPSET_ATTR_PROTOCOL_MIN as declared in ipset/ipset.h:69
	IPSET_ATTR_PROTOCOL_MIN = 10
	// IPSET_ATTR_REVISION_MIN as declared in ipset/ipset.h:70
	IPSET_ATTR_REVISION_MIN = IPSET_ATTR_PROTOCOL_MIN
	// IPSET_ATTR_INDEX as declared in ipset/ipset.h:71
	IPSET_ATTR_INDEX = 11
	// __IPSET_ATTR_CMD_MAX as declared in ipset/ipset.h:72
	__IPSET_ATTR_CMD_MAX = 12
)

const (
	// IPSET_ATTR_IP as declared in ipset/ipset.h:78
	IPSET_ATTR_IP = IPSET_ATTR_UNSPEC + 1
	// IPSET_ATTR_IP_FROM as declared in ipset/ipset.h:79
	IPSET_ATTR_IP_FROM = IPSET_ATTR_IP
	// IPSET_ATTR_IP_TO as declared in ipset/ipset.h:80
	IPSET_ATTR_IP_TO = 2
	// IPSET_ATTR_CIDR as declared in ipset/ipset.h:81
	IPSET_ATTR_CIDR = 3
	// IPSET_ATTR_PORT as declared in ipset/ipset.h:82
	IPSET_ATTR_PORT = 4
	// IPSET_ATTR_PORT_FROM as declared in ipset/ipset.h:83
	IPSET_ATTR_PORT_FROM = IPSET_ATTR_PORT
	// IPSET_ATTR_PORT_TO as declared in ipset/ipset.h:84
	IPSET_ATTR_PORT_TO = 5
	// IPSET_ATTR_TIMEOUT as declared in ipset/ipset.h:85
	IPSET_ATTR_TIMEOUT = 6
	// IPSET_ATTR_PROTO as declared in ipset/ipset.h:86
	IPSET_ATTR_PROTO = 7
	// IPSET_ATTR_CADT_FLAGS as declared in ipset/ipset.h:87
	IPSET_ATTR_CADT_FLAGS = 8
	// IPSET_ATTR_CADT_LINENO as declared in ipset/ipset.h:88
	IPSET_ATTR_CADT_LINENO = IPSET_ATTR_LINENO
	// IPSET_ATTR_MARK as declared in ipset/ipset.h:89
	IPSET_ATTR_MARK = 10
	// IPSET_ATTR_MARKMASK as declared in ipset/ipset.h:90
	IPSET_ATTR_MARKMASK = 11
	// IPSET_ATTR_CADT_MAX as declared in ipset/ipset.h:92
	IPSET_ATTR_CADT_MAX = 16
	// IPSET_ATTR_GC as declared in ipset/ipset.h:94
	IPSET_ATTR_GC = 17
	// IPSET_ATTR_HASHSIZE as declared in ipset/ipset.h:95
	IPSET_ATTR_HASHSIZE = 18
	// IPSET_ATTR_MAXELEM as declared in ipset/ipset.h:96
	IPSET_ATTR_MAXELEM = 19
	// IPSET_ATTR_NETMASK as declared in ipset/ipset.h:97
	IPSET_ATTR_NETMASK = 20
	// IPSET_ATTR_PROBES as declared in ipset/ipset.h:98
	IPSET_ATTR_PROBES = 21
	// IPSET_ATTR_RESIZE as declared in ipset/ipset.h:99
	IPSET_ATTR_RESIZE = 22
	// IPSET_ATTR_SIZE as declared in ipset/ipset.h:100
	IPSET_ATTR_SIZE = 23
	// IPSET_ATTR_ELEMENTS as declared in ipset/ipset.h:102
	IPSET_ATTR_ELEMENTS = 24
	// IPSET_ATTR_REFERENCES as declared in ipset/ipset.h:103
	IPSET_ATTR_REFERENCES = 25
	// IPSET_ATTR_MEMSIZE as declared in ipset/ipset.h:104
	IPSET_ATTR_MEMSIZE = 26
	// __IPSET_ATTR_CREATE_MAX as declared in ipset/ipset.h:106
	__IPSET_ATTR_CREATE_MAX = 27
)

const (
	// IPSET_ATTR_ETHER as declared in ipset/ipset.h:112
	IPSET_ATTR_ETHER = IPSET_ATTR_CADT_MAX + 1
	// IPSET_ATTR_NAME as declared in ipset/ipset.h:113
	IPSET_ATTR_NAME = 18
	// IPSET_ATTR_NAMEREF as declared in ipset/ipset.h:114
	IPSET_ATTR_NAMEREF = 19
	// IPSET_ATTR_IP2 as declared in ipset/ipset.h:115
	IPSET_ATTR_IP2 = 20
	// IPSET_ATTR_CIDR2 as declared in ipset/ipset.h:116
	IPSET_ATTR_CIDR2 = 21
	// IPSET_ATTR_IP2_TO as declared in ipset/ipset.h:117
	IPSET_ATTR_IP2_TO = 22
	// IPSET_ATTR_IFACE as declared in ipset/ipset.h:118
	IPSET_ATTR_IFACE = 23
	// IPSET_ATTR_BYTES as declared in ipset/ipset.h:119
	IPSET_ATTR_BYTES = 24
	// IPSET_ATTR_PACKETS as declared in ipset/ipset.h:120
	IPSET_ATTR_PACKETS = 25
	// IPSET_ATTR_COMMENT as declared in ipset/ipset.h:121
	IPSET_ATTR_COMMENT = 26
	// IPSET_ATTR_SKBMARK as declared in ipset/ipset.h:122
	IPSET_ATTR_SKBMARK = 27
	// IPSET_ATTR_SKBPRIO as declared in ipset/ipset.h:123
	IPSET_ATTR_SKBPRIO = 28
	// IPSET_ATTR_SKBQUEUE as declared in ipset/ipset.h:124
	IPSET_ATTR_SKBQUEUE = 29
	// IPSET_ATTR_PAD as declared in ipset/ipset.h:125
	IPSET_ATTR_PAD = 30
	// __IPSET_ATTR_ADT_MAX as declared in ipset/ipset.h:126
	__IPSET_ATTR_ADT_MAX = 31
)

const (
	// IPSET_ATTR_IPADDR_IPV4 as declared in ipset/ipset.h:132
	IPSET_ATTR_IPADDR_IPV4 = IPSET_ATTR_UNSPEC + 1
	// IPSET_ATTR_IPADDR_IPV6 as declared in ipset/ipset.h:133
	IPSET_ATTR_IPADDR_IPV6 = 2
	// __IPSET_ATTR_IPADDR_MAX as declared in ipset/ipset.h:134
	__IPSET_ATTR_IPADDR_MAX = 3
)

// ipset_errno as declared in ipset/ipset.h:139
type ipset_errno int32

// ipset_errno enumeration from ipset/ipset.h:139
const (
	IPSET_ERR_PRIVATE          = 4096
	IPSET_ERR_PROTOCOL         = 4097
	IPSET_ERR_FIND_TYPE        = 4098
	IPSET_ERR_MAX_SETS         = 4099
	IPSET_ERR_BUSY             = 4100
	IPSET_ERR_EXIST_SETNAME2   = 4101
	IPSET_ERR_TYPE_MISMATCH    = 4102
	IPSET_ERR_EXIST            = 4103
	IPSET_ERR_INVALID_CIDR     = 4104
	IPSET_ERR_INVALID_NETMASK  = 4105
	IPSET_ERR_INVALID_FAMILY   = 4106
	IPSET_ERR_TIMEOUT          = 4107
	IPSET_ERR_REFERENCED       = 4108
	IPSET_ERR_IPADDR_IPV4      = 4109
	IPSET_ERR_IPADDR_IPV6      = 4110
	IPSET_ERR_COUNTER          = 4111
	IPSET_ERR_COMMENT          = 4112
	IPSET_ERR_INVALID_MARKMASK = 4113
	IPSET_ERR_SKBINFO          = 4114
	IPSET_ERR_TYPE_SPECIFIC    = 4352
)

// ipset_cmd_flags as declared in ipset/ipset.h:165
type ipset_cmd_flags int32

// ipset_cmd_flags enumeration from ipset/ipset.h:165
const (
	IPSET_FLAG_BIT_EXIST                  = iota
	IPSET_FLAG_EXIST                      = (1 << IPSET_FLAG_BIT_EXIST)
	IPSET_FLAG_BIT_LIST_SETNAME           = 1
	IPSET_FLAG_LIST_SETNAME               = (1 << IPSET_FLAG_BIT_LIST_SETNAME)
	IPSET_FLAG_BIT_LIST_HEADER            = 2
	IPSET_FLAG_LIST_HEADER                = (1 << IPSET_FLAG_BIT_LIST_HEADER)
	IPSET_FLAG_BIT_SKIP_COUNTER_UPDATE    = 3
	IPSET_FLAG_SKIP_COUNTER_UPDATE        = (1 << IPSET_FLAG_BIT_SKIP_COUNTER_UPDATE)
	IPSET_FLAG_BIT_SKIP_SUBCOUNTER_UPDATE = 4
	IPSET_FLAG_SKIP_SUBCOUNTER_UPDATE     = (1 << IPSET_FLAG_BIT_SKIP_SUBCOUNTER_UPDATE)
	IPSET_FLAG_BIT_MATCH_COUNTERS         = 5
	IPSET_FLAG_MATCH_COUNTERS             = (1 << IPSET_FLAG_BIT_MATCH_COUNTERS)
	IPSET_FLAG_BIT_RETURN_NOMATCH         = 7
	IPSET_FLAG_RETURN_NOMATCH             = (1 << IPSET_FLAG_BIT_RETURN_NOMATCH)
	IPSET_FLAG_BIT_MAP_SKBMARK            = 8
	IPSET_FLAG_MAP_SKBMARK                = (1 << IPSET_FLAG_BIT_MAP_SKBMARK)
	IPSET_FLAG_BIT_MAP_SKBPRIO            = 9
	IPSET_FLAG_MAP_SKBPRIO                = (1 << IPSET_FLAG_BIT_MAP_SKBPRIO)
	IPSET_FLAG_BIT_MAP_SKBQUEUE           = 10
	IPSET_FLAG_MAP_SKBQUEUE               = (1 << IPSET_FLAG_BIT_MAP_SKBQUEUE)
	IPSET_FLAG_CMD_MAX                    = 15
)

// ipset_cadt_flags as declared in ipset/ipset.h:192
type ipset_cadt_flags int32

// ipset_cadt_flags enumeration from ipset/ipset.h:192
const (
	IPSET_FLAG_BIT_BEFORE        = iota
	IPSET_FLAG_BEFORE            = (1 << IPSET_FLAG_BIT_BEFORE)
	IPSET_FLAG_BIT_PHYSDEV       = 1
	IPSET_FLAG_PHYSDEV           = (1 << IPSET_FLAG_BIT_PHYSDEV)
	IPSET_FLAG_BIT_NOMATCH       = 2
	IPSET_FLAG_NOMATCH           = (1 << IPSET_FLAG_BIT_NOMATCH)
	IPSET_FLAG_BIT_WITH_COUNTERS = 3
	IPSET_FLAG_WITH_COUNTERS     = (1 << IPSET_FLAG_BIT_WITH_COUNTERS)
	IPSET_FLAG_BIT_WITH_COMMENT  = 4
	IPSET_FLAG_WITH_COMMENT      = (1 << IPSET_FLAG_BIT_WITH_COMMENT)
	IPSET_FLAG_BIT_WITH_FORCEADD = 5
	IPSET_FLAG_WITH_FORCEADD     = (1 << IPSET_FLAG_BIT_WITH_FORCEADD)
	IPSET_FLAG_BIT_WITH_SKBINFO  = 6
	IPSET_FLAG_WITH_SKBINFO      = (1 << IPSET_FLAG_BIT_WITH_SKBINFO)
	IPSET_FLAG_CADT_MAX          = 15
)

// ipset_create_flags as declared in ipset/ipset.h:211
type ipset_create_flags int32

// ipset_create_flags enumeration from ipset/ipset.h:211
const (
	IPSET_CREATE_FLAG_BIT_FORCEADD = iota
	IPSET_CREATE_FLAG_FORCEADD     = (1 << IPSET_CREATE_FLAG_BIT_FORCEADD)
	IPSET_CREATE_FLAG_BIT_MAX      = 7
)

// ipset_adt as declared in ipset/ipset.h:218
type ipset_adt int32

// ipset_adt enumeration from ipset/ipset.h:218
const (
	IPSET_ADD      = iota
	IPSET_DEL      = 1
	IPSET_TEST     = 2
	IPSET_ADT_MAX  = 3
	IPSET_CREATE   = IPSET_ADT_MAX
	IPSET_CADT_MAX = 4
)

// ip_set_dim as declared in ipset/ipset.h:235
type ip_set_dim int32

// ip_set_dim enumeration from ipset/ipset.h:235
const (
	IPSET_DIM_ZERO           = iota
	IPSET_DIM_ONE            = 1
	IPSET_DIM_TWO            = 2
	IPSET_DIM_THREE          = 3
	IPSET_DIM_MAX            = 6
	IPSET_BIT_RETURN_NOMATCH = 7
)

// ip_set_kopt as declared in ipset/ipset.h:249
type ip_set_kopt int32

// ip_set_kopt enumeration from ipset/ipset.h:249
const (
	IPSET_INV_MATCH      = (1 << IPSET_DIM_ZERO)
	IPSET_DIM_ONE_SRC    = (1 << IPSET_DIM_ONE)
	IPSET_DIM_TWO_SRC    = (1 << IPSET_DIM_TWO)
	IPSET_DIM_THREE_SRC  = (1 << IPSET_DIM_THREE)
	IPSET_RETURN_NOMATCH = (1 << IPSET_BIT_RETURN_NOMATCH)
)

const (
	// IPSET_COUNTER_NONE as declared in ipset/ipset.h:258
	IPSET_COUNTER_NONE = iota
	// IPSET_COUNTER_EQ as declared in ipset/ipset.h:259
	IPSET_COUNTER_EQ = 1
	// IPSET_COUNTER_NE as declared in ipset/ipset.h:260
	IPSET_COUNTER_NE = 2
	// IPSET_COUNTER_LT as declared in ipset/ipset.h:261
	IPSET_COUNTER_LT = 3
	// IPSET_COUNTER_GT as declared in ipset/ipset.h:262
	IPSET_COUNTER_GT = 4
)
