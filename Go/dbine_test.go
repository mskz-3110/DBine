package dbine

import(
  "testing"
  "github.com/stretchr/testify/assert"
  "os"
)

func TestOutputPdf(t *testing.T){
  relationshipConfig := NewRelationshipConfig()
  err := relationshipConfig.LoadJson(`{
    "Databases": [
      {
        "Name" : "Database0",
        "Tables": [
          {
            "Name": "Table0",
            "Columns": [
              {"Name": "id","Caption": "INT"},
              {"Name": "table1_id","Caption": "INT"},
              {"Name": "table2_id","Caption": "INT"},
              {"Name": "table3_id","Caption": "INT"}
            ]
          }
        ]
      },{
        "Name" : "Database1",
        "Tables": [
          {
            "Name": "Table1",
            "Columns": [
              {"Name": "id","Caption": "INT"},
              {"Name": "name","Caption": "VARCHAR"},
              {"Name": "table2_id","Caption": "INT"},
              {"Name": "table3_id","Caption": "INT"}
            ]
          },{
            "Name": "Table2",
            "Columns": [
              {"Name": "id","Caption": "INT"},
              {"Name": "name","Caption": "VARCHAR"}
            ]
          },{
            "Name": "Table3",
            "Columns": [
              {"Name": "id","Caption": "INT"},
              {"Name": "name","Caption": "VARCHAR"}
            ]
          }
        ]
      }
    ],
    "Relations": {
      "Table0.table1_id": ["Table1.id"],
      "Table0.table2_id": ["Table2.id"],
      "Table0.table3_id": ["Table3.id"],
      "Table1.table2_id": ["Table2.id"],
      "Table1.table3_id": ["Table3.id"]
    }
  }`)
  assert.NoError(t, err)

  assert.Equal(t, 2, len(relationshipConfig.Databases))
  assert.Equal(t, 1, len(relationshipConfig.Databases[0].Tables))
  assert.Equal(t, "Table0", relationshipConfig.Databases[0].Tables[0].Name)
  assert.Equal(t, 4, len(relationshipConfig.Databases[0].Tables[0].Columns))
  assert.Equal(t, "id", relationshipConfig.Databases[0].Tables[0].Columns[0].Name)
  assert.Equal(t, "INT", relationshipConfig.Databases[0].Tables[0].Columns[0].Caption)
  assert.Equal(t, "table1_id", relationshipConfig.Databases[0].Tables[0].Columns[1].Name)
  assert.Equal(t, "INT", relationshipConfig.Databases[0].Tables[0].Columns[1].Caption)
  assert.Equal(t, "table2_id", relationshipConfig.Databases[0].Tables[0].Columns[2].Name)
  assert.Equal(t, "INT", relationshipConfig.Databases[0].Tables[0].Columns[2].Caption)
  assert.Equal(t, "table3_id", relationshipConfig.Databases[0].Tables[0].Columns[3].Name)
  assert.Equal(t, "INT", relationshipConfig.Databases[0].Tables[0].Columns[3].Caption)
  assert.Equal(t, "Table1", relationshipConfig.Databases[1].Tables[0].Name)
  assert.Equal(t, 4, len(relationshipConfig.Databases[1].Tables[0].Columns))
  assert.Equal(t, "id", relationshipConfig.Databases[1].Tables[0].Columns[0].Name)
  assert.Equal(t, "INT", relationshipConfig.Databases[1].Tables[0].Columns[0].Caption)
  assert.Equal(t, "name", relationshipConfig.Databases[1].Tables[0].Columns[1].Name)
  assert.Equal(t, "VARCHAR", relationshipConfig.Databases[1].Tables[0].Columns[1].Caption)
  assert.Equal(t, "table2_id", relationshipConfig.Databases[1].Tables[0].Columns[2].Name)
  assert.Equal(t, "INT", relationshipConfig.Databases[1].Tables[0].Columns[2].Caption)
  assert.Equal(t, "table3_id", relationshipConfig.Databases[1].Tables[0].Columns[3].Name)
  assert.Equal(t, "INT", relationshipConfig.Databases[1].Tables[0].Columns[3].Caption)
  assert.Equal(t, "Table2", relationshipConfig.Databases[1].Tables[1].Name)
  assert.Equal(t, 2, len(relationshipConfig.Databases[1].Tables[1].Columns))
  assert.Equal(t, "id", relationshipConfig.Databases[1].Tables[1].Columns[0].Name)
  assert.Equal(t, "INT", relationshipConfig.Databases[1].Tables[1].Columns[0].Caption)
  assert.Equal(t, "name", relationshipConfig.Databases[1].Tables[1].Columns[1].Name)
  assert.Equal(t, "VARCHAR", relationshipConfig.Databases[1].Tables[1].Columns[1].Caption)
  assert.Equal(t, "Table3", relationshipConfig.Databases[1].Tables[2].Name)
  assert.Equal(t, 2, len(relationshipConfig.Databases[1].Tables[2].Columns))
  assert.Equal(t, "id", relationshipConfig.Databases[1].Tables[2].Columns[0].Name)
  assert.Equal(t, "INT", relationshipConfig.Databases[1].Tables[2].Columns[0].Caption)
  assert.Equal(t, "name", relationshipConfig.Databases[1].Tables[2].Columns[1].Name)
  assert.Equal(t, "VARCHAR", relationshipConfig.Databases[1].Tables[2].Columns[1].Caption)

  assert.Equal(t, 5, len(relationshipConfig.Relations))
  assert.Equal(t, map[string][]string{
    "Table0.table1_id": {"Table1.id"},
    "Table0.table2_id": {"Table2.id"},
    "Table0.table3_id": {"Table3.id"},
    "Table1.table2_id": {"Table2.id"},
    "Table1.table3_id": {"Table3.id"},
  }, relationshipConfig.Relations)

  dotString := relationshipConfig.ToDotString(DefaultDotStyle())
  assert.True(t, 0 < len(dotString))
  dotFilePath := "../Dot/Sample.dot"
  _, err = os.Stat(dotFilePath)
  if os.IsNotExist(err) {
    err = os.WriteFile(dotFilePath, []byte(dotString), 0644)
    assert.NoError(t, err)
  } else {
    readBuffer, err := os.ReadFile(dotFilePath)
    assert.NoError(t, err)
    assert.Equal(t, string(readBuffer), dotString)
  }

  pdfBuffer := DotToPdfBuffer(dotString)
  assert.True(t, 0 < len(pdfBuffer))
  os.WriteFile("../Dot/Sample.pdf", pdfBuffer, 0644)
}
