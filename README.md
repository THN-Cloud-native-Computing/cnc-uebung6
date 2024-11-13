# Übung 6 - Microservice Architekturen

![microservices.jpg](microservices.jpg)

**Hinweise:**

Zunächst einmal können Sie dieses Repository wieder über git klonen.  

Legen Sie für jede der Aufgaben einen Ordner an.   

In dieser Übung werden wir die Programmiersprache Go verwenden. Diese ist im cloud-nativen Umfeld sehr beliebt. Kernkomponenten von Docker und Kubernetes sind bspw. in Go geschrieben.  

Installieren Sie zunächst [Go](https://go.dev/)  

Ob die Installation erfolgreich war, können Sie testen durch: 

   ```bash
go version
   ```
Wichtige Befehle können Sie sich anzeigen lassen mit:

   ```bash
go help
   ```

Hier können Sie sich zunächst einen Überblick über die wichtigsten Sprachkonstrukte verschaffen: [https://www.golang-book.com/books/intro](https://www.golang-book.com/books/intro)  

Hier finden Sie die Standard-Library: [https://pkg.go.dev/std](https://pkg.go.dev/std)  

Außerdem werden wir in dieser Übung Postman verwenden, womit man u.a. APIs testen kann. Sie können das Tool hier kostenlos herunterladen bzw. sich kostenlos registrieren: [https://www.postman.com/](https://www.postman.com/)

**Aufgabe 1 - Eine vereinfachte Rest API mit Go entwickeln und mit Postman testen**

Erstellen Sie ein Verzeichnis für diese Aufgabe und legen Sie die Datei main.go darin ab.  

Erklären Sie, was das Programm tut.  

Öffnen Sie nun ein Terminal und navigieren Sie zum Verzeichnis für diese Aufgabe.  

Starten Sie das Programm mit:

   ```bash
go run main.go
   ```

Rufen Sie in ihrem Browser nun

   ```bash
http://localhost:8080
   ```
auf, und vergewissern sie sich, dass der Server Requests korrekt beantwortet.  

Testen Sie nun mit Postman, wie der Server auf GET, PUT, POST und DELETE Requests reagiert.

**Aufgabe 2 - Anpassung der API zum Handling unterschiedlicher http-Requests**

Passen Sie das Programm aus Aufgabe 1 so an, dass auf die unterschiedlichen http-Requests (GET, PUT, POST und DELETE) auch unterschiedliche Responses erfolgen.  

Hinweis: In der Funktion helloHandler kann durch die Abfrage ob `r.Method` einer `http.MethodGet` entspricht bspw. überprüft werden, ob es sich um einen GET-Request handelt.  

Prüfen Sie nun mit Postman, ob ihr Programm auf unterschiedliche Requests auch entsprechend reagiert.  

**Aufgabe 3 - HTML-Responses**

Passen Sie das Programm aus Aufgabe 2 so an, dass statt reinem Text, html-Code zurückgegeben wird.  

Hinweis: In der Funktion `Fprintf` kann statt Text auch in backticks gewrappter html-Code eingefügt werden.    

Prüfen Sie mit ihrem Browser bzw. mit Postman, ob ihr Programm auf unterschiedliche Requests auch entsprechenden html-Code zurück gibt. 

**Aufgabe 3 - Zwei einfache kommunizierende Microservices**

Im Verzeichnis 2-simple-microservices finden Sie in den Unterverzeichnissen zwei einfache in Go implementierte Microservices. Starten Sie sowohl im Verzeichnis service-a als auch im Verzeichnis service-b die beiden Dienste mit

   ```bash
go run main.go
   ```

Öffnen Sie ihren Browser mit der Adresse

   ```bash
http://localhost:8080/service-a
   ```
und prüfen Sie, ob die Kommunikation der beiden Dienste funktioniert.  

Wie sind die beiden Dienste gekoppelt? Ist das gut? Diskutieren Sie das!

**Aufgabe 2 - Ein erstes Kubernetes Cluster aufsetzen**

Installieren Sie [kubectl](https://kubernetes.io/de/docs/tasks/tools/install-kubectl/)

Überprüfen Sie über die Kommandozeile, ob kubectl erfolgreich installiert wurde mittels
   ```bash
kubectl version
   ```
Sie können nun ein einfaches Test-Cluster mit einem einfachen Server starten:

   ```bash
kubectl create deployment hello-node --image=registry.k8s.io/e2e-test-images/agnhost:2.39 -- /agnhost netexec --http-port=8080

   ```
Lassen Sie sich die laufenden Deployments anzeigen mit:
   ```bash
kubectl get deployments

   ```
Lassen Sie sich die laufenden Pods anzeigen mit:
   ```bash
kubectl get pods

   ```
Öffnen Sie das Docker Desktop Dashboard und stoppen Sie dort den laufenden Pod bzw. Container. Was passiert?

Lassen Sie sich in der Kommandozeile die Cluster Events anzeigen:
   ```bash
kubectl get events

   ```
Öffnen Sie ihren Webbrowser mit 
   ```bash
http://localhost:8080
   ```
Was ist zu sehen?

Fügen Sie dem Cluster einen Load Balancer Service hinzu, der den Dienst nach außen hin verfügbar macht:
   ```bash
kubectl expose deployment hello-node --type=LoadBalancer --port=8080

   ```
Prüfen Sie, ob der Service läuft mit:
   ```bash
kubectl get services

   ```
Probieren Sie nun noch einmal den Aufruf über ihren Webbrowser.

Löschen Sie die Ressourcen anschließend wieder mit:
   ```bash
kubectl delete service hello-node

   ```
   ```bash
kubectl delete deployment hello-node

   ```
**Aufgabe 3 - Ein Kubernetes Cluster mit mehreren Pods starten**

Mit der Konfigurationsdatei 
   ```bash
nginx-deployment.yaml

   ```
können Sie ein einfaches Kubernetes Cluster starten, in dem ein nginx Server in einem Pod läuft und über den NodePort Service von außen verfügbar macht.  

Starten Sie das Cluster einmal mit
   ```bash
kubectl apply -f nginx-deployment.yaml

   ```
Sie sollten den Dienst in ihrem Browser dann erreichen können über:
   ```bash
http://localhost:30000

   ```
Löschen Sie die Ressourcen anschließend wieder.  

Ändern Sie die Konfigurationsdatei so, dass jeweils 3 Pods mit Servern gestartet werden. Was passiert, wenn Sie die Container stoppen?

**Aufgabe 4 - Resource Limits**

Passen Sie die Konfigurationsdatei aus Aufgabe 3 zunächst wieder so an, dass nur ein Pod gestartet wird.  

Ändern Sie die Konfigurationsdatei nun so, dass für die CPU- und Speicher-Nutzung des Containers folgende Mindest- und Höchstgrenzen gesetzt werden:

- Untergrenzen: CPU 250m / RAM 64Mi
- Obergrenzen: CPU 500m / RAM 128Mi

Was bedeuten die Angaben?  

- Starten Sie das Cluster und schauen Sie sich unter den "Stats" im Docker Desktop Dashboard die CPU- und RAM-Nutzung an.  

- Rufen Sie den Server über ihren Webbrowser auf und schauen Sie sich die Änderungen in der Auslastung in "Stats" an.

- Was würde passieren, wenn die Obergrenzen für CPU und RAM überschritten werden?  

- Finden Sie heraus, welches die kleinstmöglichen Grenzen für CPU und Speicher sind. Passen Sie die Konfigurationsdatei entsprechend an.   

- Starten Sie das Cluster erneut und prüfen Sie im Docker Desktop Dashboard den Status des Containers.  

- Experimentieren Sie mit anderen Werten für die Ressource Limits und prüfen Sie den Status des Containers.

**Aufgabe 5 - Automatisches Skalieren mit den Horizontal Pod Autoscaler (HPA)**

Setzen Sie die Ressource Limits zunächst wieder so, wie zu Beginn von Aufgabe 4.  

Konfigurieren Sie nun einen Horizontal Pod Autoscaler (HPA), der bei überschreiten von 80% der CPU-Auslastung einen weiteren Pod startet.

Rufen Sie den Server über ihren Webbrowser auf und schauen Sie sich die Änderungen in der Auslastung in "Stats" an.

Installieren Sie sich das Programm [GNU Parallels](https://www.gnu.org/software/parallel/)  

Hinweis: Sie können das Programm über Homebrew einfach installieren über:
   ```bash
brew install parallel

   ```

Über folgenden Befehl können Sie eine Sequenz von 1000 parallelen Request (jeweils 100) an den Server schicken:
   ```bash
seq 1 1000 | parallel -j 100 curl -s http://localhost:30000/

   ```
Schauen Sie sich nun die Änderungen in der Auslastung in "Stats" an.  

Experimentieren Sie mit den Ressource Limits und der Anzahl und dem Umfang der parallelen Request so, dass man die Aktivitäten des HPA im Dashboard verfolgen kann.  

![cpu-usage.png](cpu-usage.png)

**Aufgabe 6 - Health Checks**

Nehmen Sie die ursprünglich Konfigurationsdatei zu Aufgabe 3 als Ausgangspunkt.  Die Konfiguration soll nun um Health Checks erweitert werden.  

Was bedeuten in diesem Zusammenhang die Begriffe Liveness Probe und Readiness Probe?  

Passen Sie die Konfigurationsdatei nun so an, dass folgende Health Checks ausgeführt werden:

- Liveness Probe initial nach 15 Sekunden, danach alle 20 Sekunden
- Readiness Probe initial nach 10 Sekunden, danach alle 10 Sekunden

Prüfen Sie im Docker Desktop Dashboard unter "Logs" die Ausführung der Health Checks.

 
