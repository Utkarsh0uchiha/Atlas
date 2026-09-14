console.log("Atlas dashboard JS loaded");
fetch("/api/status")
    .then(response => response.json())
    .then(data => {

        data.backends.forEach(backendData => {
            const backendElement = document.getElementById(`backend-${backendData.id}`);
            const status = backendElement.querySelector(".status")
            const statusText = status.querySelector(".status-text")
            const requests = backendElement.querySelector(".requests")
            const latency = backendElement.querySelector(".latency")
            if(backendData.healthy){
                statusText.textContent = "Healthy";
                status.classList.remove("unhealthy");
                status.classList.add("healthy");
            } else {
                statusText.textContent = "Unhealthy";
                status.classList.remove("healthy");
                status.classList.add("unhealthy");
            }

            requests.textContent = backendData.requests
            latency.textContent = backendData.average_latency
        });
    });
