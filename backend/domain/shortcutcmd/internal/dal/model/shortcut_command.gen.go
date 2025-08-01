// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/coze-dev/coze-studio/backend/api/model/ocean/cloud/playground"

const TableNameShortcutCommand = "shortcut_command"

// ShortcutCommand bot shortcut command table
type ShortcutCommand struct {
	ID              int64                        `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`                                                             // id
	ObjectID        int64                        `gorm:"column:object_id;not null;comment:Entity ID, this command can be used for this entity" json:"object_id"`                   // Entity ID, this command can be used for this entity
	CommandID       int64                        `gorm:"column:command_id;not null;comment:command id" json:"command_id"`                                                          // command id
	CommandName     string                       `gorm:"column:command_name;not null;comment:command name" json:"command_name"`                                                    // command name
	ShortcutCommand string                       `gorm:"column:shortcut_command;not null;comment:shortcut command" json:"shortcut_command"`                                        // shortcut command
	Description     string                       `gorm:"column:description;not null;comment:description" json:"description"`                                                       // description
	SendType        int32                        `gorm:"column:send_type;not null;comment:send type 0:query 1:panel" json:"send_type"`                                             // send type 0:query 1:panel
	ToolType        int32                        `gorm:"column:tool_type;not null;comment:Type 1 of tool used: WorkFlow 2: Plugin" json:"tool_type"`                               // Type 1 of tool used: WorkFlow 2: Plugin
	WorkFlowID      int64                        `gorm:"column:work_flow_id;not null;comment:workflow id" json:"work_flow_id"`                                                     // workflow id
	PluginID        int64                        `gorm:"column:plugin_id;not null;comment:plugin id" json:"plugin_id"`                                                             // plugin id
	PluginToolName  string                       `gorm:"column:plugin_tool_name;not null;comment:plugin tool name" json:"plugin_tool_name"`                                        // plugin tool name
	TemplateQuery   string                       `gorm:"column:template_query;comment:template query" json:"template_query"`                                                       // template query
	Components      []*playground.Components     `gorm:"column:components;comment:Panel parameters;serializer:json" json:"components"`                                             // Panel parameters
	CardSchema      string                       `gorm:"column:card_schema;comment:card schema" json:"card_schema"`                                                                // card schema
	ToolInfo        *playground.ToolInfo         `gorm:"column:tool_info;comment:Tool information includes name+variable list;serializer:json" json:"tool_info"`                   // Tool information includes name+variable list
	Status          int32                        `gorm:"column:status;not null;comment:Status, 0 is invalid, 1 is valid" json:"status"`                                            // Status, 0 is invalid, 1 is valid
	CreatorID       int64                        `gorm:"column:creator_id;comment:creator id" json:"creator_id"`                                                                   // creator id
	IsOnline        int32                        `gorm:"column:is_online;not null;comment:Is online information: 0 draft 1 online" json:"is_online"`                               // Is online information: 0 draft 1 online
	CreatedAt       int64                        `gorm:"column:created_at;not null;autoCreateTime:milli;comment:Create Time in Milliseconds" json:"created_at"`                    // Create Time in Milliseconds
	UpdatedAt       int64                        `gorm:"column:updated_at;not null;autoUpdateTime:milli;comment:Update Time in Milliseconds" json:"updated_at"`                    // Update Time in Milliseconds
	AgentID         int64                        `gorm:"column:agent_id;not null;comment:When executing a multi instruction, which node executes the instruction" json:"agent_id"` // When executing a multi instruction, which node executes the instruction
	ShortcutIcon    *playground.ShortcutFileInfo `gorm:"column:shortcut_icon;comment:shortcut icon;serializer:json" json:"shortcut_icon"`                                          // shortcut icon
	PluginToolID    int64                        `gorm:"column:plugin_tool_id;not null;comment:tool_id" json:"plugin_tool_id"`                                                     // tool_id
}

// TableName ShortcutCommand's table name
func (*ShortcutCommand) TableName() string {
	return TableNameShortcutCommand
}
