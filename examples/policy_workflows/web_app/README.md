# Web App Microsegmentation Policy Workflow  

This example defines a very simple web application in the Illumio PCE. The two-tier application is comprised of Tomcat web servers and a MySQL database. In the production environment, the application is load balanced between datacenters in US-East and US-West zones.  

A single ruleset is defined to set up microsegmentation for communication between the web app workloads. The Tomcat servers communicate over the default MySQL port, 3306, with the databases. Web traffic on ports 80 and 443 are permitted from all IPs to the production web servers.  

The approach used here can be extended to more complex application architectures and shows how easily microsegmentation can be configured in Illumio.  
