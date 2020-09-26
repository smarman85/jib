package newPipe

import (
  f "fmt"
  "encoding/json"
)

type pipeline struct {
        Schema string `json:"schema"`
        Id string `json:"id"`
        Protect bool `json:"protect"`
        Variables []vars `json:"variables"`
        Metadata meta `json:"metadata"`
        Pipeline pline `json:"pipeline"`
}

type vars struct {
        Type string `json:"type"`
        Default string `json:"defaultValue"`
        Desc string `json:"description"`
        Name string `json:"name"`
}

type meta struct {
        Name string `json:"name"`
        Desc string `json:"description"`
        Owner string `json:"owner"`
        Scopes []string `json:"scopes"`
}

type pline struct {
        Modified string `json:"lastModifiedBy"`
        Ts string `json:"updateTs"`
        Params []string `json:"parameterConfig"`
        limitConcurrent bool `json:"limitConcurrent"`
        keepWaiting bool `json:"keepWaitingPipelines"`
        Desc string `json:"description"`
        Triggers []string `json:"triggers"`
        Notification []string `json:"notifications"`
        Stages []string `json:"stages"`
}


func Hello(args []string) {
  //f.Println("Hello from newPipeline")
  pipe := pipeline{
          Schema: "v2",
          Id: args[0],
          Protect: false,
          Variables: []vars{
                  vars{
                      Type: "int",
                      Default: "22",
                      Desc: "This is a blah, blah...",
                      Name: args[0],
                  },
          },
          Metadata: meta{
                      Name: args[0],
                      Desc: "Somejunk",
                      Owner: "me@mail.com",
                      Scopes: []string{"global"},
          },
          Pipeline: pline{
                  Modified: "",
                  Ts: "0",
                  Params: []string{},
                  limitConcurrent: false,
                  keepWaiting: false,
                  Desc: "More info ...",
                  Triggers: []string{},
                  Notification: []string{},
                  Stages: []string{},
          },
  }
  p, err := json.MarshalIndent(&pipe, "", "    ")
  if err != nil {
          f.Println("Error marshalling pipeline", err)
  }
  f.Println(string(p))
}
