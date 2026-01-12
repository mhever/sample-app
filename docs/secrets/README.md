Instead of storing the secrets in git, right now they are created in k8s manually:

~~~
apiVersion: v1
kind: Secret
metadata:
  name: sample-app-secret
type: Opaque
data:
  DB_PASSWORD: *base64encodedstring*
~~~

becomes
~~~
sudo k3s kubectl create secret generic postgres-secret --from-literal=postgres-password='...' --dry-run=client -o yaml | sudo k3s kubectl apply -f -
sudo k3s kubectl create secret generic sample-app-secret --from-literal=DB_PASSWORD='...' --dry-run=client -o yaml | sudo k3s kubectl apply -f -
~~~