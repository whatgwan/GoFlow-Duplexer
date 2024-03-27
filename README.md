# GoFlow-Duplexer
a Go powered netflow packet receiver and sender. Receive one netflow stream, and send it to as many destinations as required


**Prerequisites**



Before using the NetFlow Duplication Go Application, ensure that you have:


A Go environment set up on your system so you can run the code or build a binary file.



**Starting the Application:**



Once compiled, you can start the application as you would any other script or compiled go-binary, just run:



"/path/to/goflow-duplexer" with the required privilidges with the following args:


-devices="device1_ip:port1,device2_ip:port2" (the devices and ports you wish to send netflow to) 



-port=2055 (the port you are receiving netflow traffic on)



**Example:**



'nohup /path/to/netflow-duplication-app -devices="device1_ip:port1,device2_ip:port2" -port=2055 >/dev/null 2>&1 &'


To start the application in the background



**This application is distributed under the MIT License.**
 
