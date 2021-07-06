{{- define "service" }}
{{- $name := index . 0 }}

apiVersion: v1
kind: Service
metadata:
  name: {{ $name }}
spec:
  selector:
    app: {{ $name }}
  ports:
    - port: 80
      targetPort: 8080
      protocol: TCP
{{- end }}