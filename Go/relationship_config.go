package dbine

import(
  "encoding/json"
  "fmt"
  "strings"
  "slices"
)

type RelationshipColumn struct{
  Name string
  Caption string
}

type RelationshipTable struct{
  Name string
  Columns []RelationshipColumn
}

type RelationshipDatabase struct{
  Name string
  Tables []RelationshipTable
}

type RelationshipConfig struct{
  Databases []RelationshipDatabase
  Relations map[string][]string
}

type DotStyle struct{
  FontName string
  FontColor string
  RelationColor string
  BackgroundColor string
  Dpi int
}

func DefaultDotStyle()(DotStyle){
  return DotStyle{
    FontName: "Yu Mincho Demibold",
    FontColor: "black",
    RelationColor: "gray",
    BackgroundColor: "white",
    Dpi: 350,
  }
}

func NewRelationshipConfig()(*RelationshipConfig){
  return &RelationshipConfig{}
}

func (relationshipConfig *RelationshipConfig)LoadJson(relationshipConfigJson string)(error){
  return json.Unmarshal([]byte(relationshipConfigJson), &relationshipConfig)
}

func (relationshipConfig RelationshipConfig)ToDotString(dotStyle DotStyle)(string){
  var sb strings.Builder
  sb.WriteString(fmt.Sprintf(`digraph {
  graph [
    fontname = "%s"
    fontcolor = "%s"
    color = "%s"
    bgcolor = "%s"
    dpi = %d
    margin = -0.01
    rankdir = LR
  ]
  node [
    fontname = "%s"
    fontcolor = "%s"
    shape = none
  ]
  edge [
    fontname = "%s"
    fontcolor = "%s"
    color = "%s:%s:%s"
    dir = none
  ]
  
`, dotStyle.FontName, dotStyle.FontColor, dotStyle.FontColor, dotStyle.BackgroundColor, dotStyle.Dpi, dotStyle.FontName, dotStyle.FontColor, dotStyle.FontName, dotStyle.FontColor, dotStyle.RelationColor, dotStyle.BackgroundColor, dotStyle.RelationColor))
  for _, database := range relationshipConfig.Databases {
    sb.WriteString(fmt.Sprintf(`  subgraph cluster_%s {
    label = "%s"
    labeljust = l
  
`, database.Name, database.Name))
    for _, table := range database.Tables {
      sb.WriteString(fmt.Sprintf(`    %s [label = <<table border="1" cellspacing="0" cellpadding="0" color="%s" bgcolor="%s">
      <tr><td colspan="2"><b>%s</b></td></tr>
`, table.Name, dotStyle.RelationColor, dotStyle.RelationColor, table.Name))
      for _, column := range table.Columns {
        sb.WriteString(fmt.Sprintf(`      <tr><td bgcolor="%s" cellpadding="2" port="%s"> %s </td><td bgcolor="%s" cellpadding="2" align="left"><i>%s</i> </td></tr>`, dotStyle.BackgroundColor, column.Name, column.Name, dotStyle.BackgroundColor, column.Caption))
        sb.WriteString("\n")
      }
      sb.WriteString("    </table>>]\n")
    }
    sb.WriteString("  }\n")
  }
  relationKeys := []string{}
  for key, _ := range relationshipConfig.Relations {
    relationKeys = append(relationKeys, key)
  }
  slices.Sort(relationKeys)
  for _, key := range relationKeys {
    replacedKey := strings.Replace(key, ".", ":", -1)
    relationValues := relationshipConfig.Relations[key]
    slices.Sort(relationValues)
    for _, value := range relationValues {
      sb.WriteString(fmt.Sprintf("  %s -> %s\n", replacedKey, strings.Replace(value, ".", ":", -1)))
    }
  }
  sb.WriteString("}")
  return sb.String()
}

func (relationshipConfig RelationshipConfig)ToPdfBuffer(dotStyle DotStyle)([]byte){
  return DotToPdfBuffer(relationshipConfig.ToDotString(dotStyle))
}
