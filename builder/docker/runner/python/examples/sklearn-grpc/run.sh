docker run --name python-model-runner --env pipecoupler_port=5000 --env PythonModel.plugin=artifacts.predictor --env PythonModel.URI=/iris-classifier  -v /Users/steven/Applications/Dev-Ops/src/github.com/TIBCOSoftware/python/docker/sklearn/artifacts:/app/artifacts -p 5000:5000 bigoyang/python-model:0.1.0