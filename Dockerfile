FROM golang:1.12 AS build1
COPY http-api.go /
RUN cd / && go build http-api.go

FROM python:3
RUN pip install numpy Pillow opencv-python
COPY . /buildspace
RUN cd /buildspace && python setup.py install

EXPOSE 8080
WORKDIR /workspace
COPY --from=build1 /http-api /workspace/
RUN cp /buildspace/image_to_text.py .
CMD ["/workspace/http-api", "server"]
