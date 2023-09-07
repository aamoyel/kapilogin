# kapilogin

<!-- TABLE OF CONTENTS -->
<details open>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">Purpose</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#configuration">Configuration</a></li>
    <li><a href="#Contribute">Contribute</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>
</br>


<!-- ABOUT THE PROJECT -->
## Purpose
This project allows you to dynamicaly retrive kubeconfig files and use kubelogin for oidc login to authenticate on clusters managed by Cluster API.  
<p align="right">(<a href="#top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

### Prerequisites
You need to have :
* [Kubelogin](https://github.com/int128/kubelogin) installed on your machine.
* A Kubernetes cluster with Cluster API and child clusters bootstraped with it.
* Your cluster can assign IPs on Services type LoadBalancer.
* kubectl binary

### Installation
1. Deploy the latest kapilogin server release on the Kubernetes with Cluster API :
   ```sh
    kubectl kustomize https://github.com/aamoyel/kapilogin/deploy | kubectl apply -f -
   ```
2. Get the latest release of the CLI and add it in your PATH

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONFIGURATION -->
## Configuration
1. First, you need to get the LoadBalancer IP use by kapilogin API:
    ```sh
     kubectl -n kapilogin get svc kapilogin -o json | jq '.status.loadBalancer.ingress[0].ip'
    ```
2. To authenticate on your clusters and define Kapilogin API endpoint, you need to configure Kapilogin. You can use an url to the raw file (eg: https://raw.githubusercontent.com/project/main/kapilogin.yaml) or directly create the file on you system with the command below:
    ```sh
    cat <<EOF > $HOME/.kapilogin.yaml
    kapiloginApiEndpoint: KAPILOGIN_API_ENDPOINT # LoadBalancer IP
    oidcIssuerUrl: ISSUER_URL
    oidcClientId: YOUR_CLIENT_ID
    oidcClientSecret: YOUR_CLIENT_SECRET # Optional
    EOF
    ```
3. To use this configuration you can pass "-c CFG_PATH" to the kapilogin CLI or set the var KAPILOGIN_CONFIG=... (url of local file path)

4. Now, you can use the 'kapilogin' CLI.
    ```sh
     kapilogin --help
    ```

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- Contribute -->
## Contribute
You can create issues and PRs on this project if you have any problems or suggestions.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the Apache-2.0 license. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Alan Amoyel - [@AlanAmoyel](https://twitter.com/AlanAmoyel)

<p align="right">(<a href="#top">back to top</a>)</p>
