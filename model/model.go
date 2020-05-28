//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"time"
)

type ColumnsAgent struct {
	ID, Hash, UserAgent, Platform, Browser, Version, CreatedAt string
}

type ColumnsIp struct {
	ID, Ip, CountryCode, CountryName, CreatedAt string
}

type ColumnsLink struct {
	ID, TypeID, Url, TargetUrl, Viewed, CreatedAt, UpdatedAt string
}

type ColumnsLog struct {
	ID, LinkID, IpID, AgentID, Others, CreatedAt string
}

type ColumnsType struct {
	ID, Name, Status, CreatedAt string
}

type ColumnsSt struct {
	Agent ColumnsAgent
	Ip    ColumnsIp
	Link  ColumnsLink
	Log   ColumnsLog
	Type  ColumnsType
}

var Columns = ColumnsSt{
	Agent: ColumnsAgent{
		ID:        "id",
		Hash:      "hash",
		UserAgent: "user_agent",
		Platform:  "platform",
		Browser:   "browser",
		Version:   "version",
		CreatedAt: "created_at",
	},
	Ip: ColumnsIp{
		ID:          "id",
		Ip:          "ip",
		CountryCode: "country_code",
		CountryName: "country_name",
		CreatedAt:   "created_at",
	},
	Link: ColumnsLink{
		ID:        "id",
		TypeID:    "type_id",
		Url:       "url",
		TargetUrl: "target_url",
		Viewed:    "viewed",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	},
	Log: ColumnsLog{
		ID:        "id",
		LinkID:    "link_id",
		IpID:      "ip_id",
		AgentID:   "agent_id",
		Others:    "others",
		CreatedAt: "created_at",
	},
	Type: ColumnsType{
		ID:        "id",
		Name:      "name",
		Status:    "status",
		CreatedAt: "created_at",
	},
}

type TableAgent struct {
	Name, Alias string
}

type TableIp struct {
	Name, Alias string
}

type TableLink struct {
	Name, Alias string
}

type TableLog struct {
	Name, Alias string
}

type TableType struct {
	Name, Alias string
}

type TablesSt struct {
	Agent TableAgent
	Ip    TableIp
	Link  TableLink
	Log   TableLog
	Type  TableType
}

var Tables = TablesSt{
	Agent: TableAgent{
		Name:  "agents",
		Alias: "t",
	},
	Ip: TableIp{
		Name:  "ips",
		Alias: "t",
	},
	Link: TableLink{
		Name:  "links",
		Alias: "t",
	},
	Log: TableLog{
		Name:  "logs",
		Alias: "t",
	},
	Type: TableType{
		Name:  "types",
		Alias: "t",
	},
}

type Agent struct {
	tableName struct{} `sql:"agents,alias:t" pg:",discard_unknown_columns"`

	ID        int        `sql:"id,pk"`
	Hash      string     `sql:"hash,notnull"`
	UserAgent string     `sql:"user_agent,notnull"`
	Platform  *string    `sql:"platform"`
	Browser   *string    `sql:"browser"`
	Version   *string    `sql:"version"`
	CreatedAt *time.Time `sql:"created_at"`
}

type Ip struct {
	tableName struct{} `sql:"ips,alias:t" pg:",discard_unknown_columns"`

	ID          int        `sql:"id,pk"`
	Ip          string     `sql:"ip,notnull"`
	CountryCode *string    `sql:"country_code"`
	CountryName *string    `sql:"country_name"`
	CreatedAt   *time.Time `sql:"created_at"`
}

type Link struct {
	tableName struct{} `sql:"links,alias:t" pg:",discard_unknown_columns"`

	ID        int        `sql:"id,pk"`
	TypeID    int        `sql:"type_id,notnull"`
	Url       string     `sql:"url,notnull"`
	TargetUrl string     `sql:"target_url,notnull"`
	Viewed    *int       `sql:"viewed"`
	CreatedAt *time.Time `sql:"created_at"`
	UpdatedAt *time.Time `sql:"updated_at"`
}

type Log struct {
	tableName struct{} `sql:"logs,alias:t" pg:",discard_unknown_columns"`

	ID        int        `sql:"id,pk"`
	LinkID    int        `sql:"link_id,notnull"`
	IpID      int        `sql:"ip_id,notnull"`
	AgentID   int        `sql:"agent_id,notnull"`
	Others    *string    `sql:"others"`
	CreatedAt *time.Time `sql:"created_at"`
}

type Type struct {
	tableName struct{} `sql:"types,alias:t" pg:",discard_unknown_columns"`

	ID        int        `sql:"id,pk"`
	Name      string     `sql:"name,notnull"`
	Status    *int       `sql:"status"`
	CreatedAt *time.Time `sql:"created_at"`
}
