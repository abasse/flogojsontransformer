# JSON transformer flogo activity
This activity allows your flogo application to transform JSON data. The transformation is done with kazaam (https://github.com/qntfy/kazaam).



## Installation

```bash
flogo install github.com/abasse/flogojsontransformer
```

## Schema
Inputs and Outputs:

```json
  { "inputs":[
        {
          "name": "Input",
          "type": "string",
          "required": true
        },
        {
          "name": "Transformation",
          "type": "string",
          "required": true
        }
      ],
      "outputs": [
         {
           "name": "result",
           "type": "string"
          }
      ]
 }
```
