import tr
import json
import sys

data = tr.run_angle(sys.argv[1])
dataList = []

for i, val in enumerate(data):
    dataList.append( {
        "cx" : float(val[0][0]),
    	"cy" : float(val[0][1]),
    	"width" : float(val[0][2]),
    	"height": float(val[0][3]),
    	"angle": float(val[0][4]),
    	"text" : str(val[1]),
    	"confidence": float(val[2]),
    })

print(json.dumps(dataList))

