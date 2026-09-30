console.log("Atlas dashboard JS loaded");
let trafficInterval = null

fetch("/api/status")
    .then(response => response.json())
    .then(data => {
        
        updateOverview(data)
        const startTraffic = document.querySelector(".traffic-start")
        const stopTraffic = document.querySelector(".traffic-stop")
        
        startTraffic.addEventListener("click", () => {
            if (trafficInterval !== null){
                return 
            }

            trafficInterval = setInterval(() => {
                fetch("/")
                .catch(error => console.error("Traffic request failed:", error))
            }, 1000)

            startTraffic.disabled = true
            stopTraffic.disabled = false
        })

        

        stopTraffic.addEventListener("click", () => {
            if(trafficInterval === null){
                return
            }

            clearInterval(trafficInterval)
            trafficInterval = null

            startTraffic.disabled = false
            stopTraffic.disabled = true
        })
        
        startTraffic.disabled = false
        stopTraffic.disabled = true

        data.backends.forEach(backendData => {
            const backendElement = document.getElementById(`backend-${backendData.id}`);
            const requests = backendElement.querySelector(".requests")
            const latency = backendElement.querySelector(".latency")
            const toggle = backendElement.querySelector(".backend-toggle")
            
           updateBackendHealth(backendData)

            if(backendData.enabled){
                        toggle.textContent = "Disable";
                        toggle.classList.remove("enable");
                        toggle.classList.add("disable");
            } else {
                        toggle.textContent = "Enable";
                        toggle.classList.remove("disable");
                        toggle.classList.add("enable");
            }

            toggle.addEventListener("click", () => {
                const action = backendData.enabled ? "disable" : "enable"
                const url = `/api/backends/${backendData.id}/${action}`
                toggle.disabled = true
                fetch(url, {
                    method: "POST"
                })
                .then(response => {
                    if(!response.ok){
                        throw new Error("Failed to update backend")
                    }
                    return response.json()
                })
                .then(data => {
                    backendData.enabled = !backendData.enabled
                    if(backendData.enabled){
                        toggle.textContent = "Disable";
                        toggle.classList.remove("enable");
                        toggle.classList.add("disable");
                    } else {
                        toggle.textContent = "Enable";
                        toggle.classList.remove("disable");
                        toggle.classList.add("enable");
                    }
                })
                .catch(error => {
                    console.error(error)
                })
                .finally(() => {
                    toggle.disabled = false
                })
            })

            requests.textContent = backendData.requests;
            latency.textContent = backendData.average_latency;
        });
    });

function updateBackendHealth(backendData){
    const backendElement = document.getElementById(`backend-${backendData.id}`);
    const status = backendElement.querySelector(".status")
    const statusText = status.querySelector(".status-text")
     if(backendData.healthy){
                statusText.textContent = "Healthy";
                status.classList.remove("unhealthy");
                status.classList.add("healthy");
            } else {
                statusText.textContent = "Unhealthy";
                status.classList.remove("healthy");
                status.classList.add("unhealthy");
            }
}

function updateOverview(data){
        const totalRequests = document.querySelector(".total-requests")
        const healthyBackends = document.querySelector(".healthy-backends")
        const uptime = document.querySelector(".uptime")
        
        totalRequests.textContent = data.total_requests
        healthyBackends.textContent = `${data.healthy_backends}/${data.total_backends}`
        uptime.textContent = data.uptime
}

function updateBackendMetrics(backendData) {
    const backendElement = document.getElementById(`backend-${backendData.id}`)

    backendElement.querySelector(".requests").textContent = backendData.requests
    backendElement.querySelector(".latency").textContent = backendData.average_latency
}

const events = new EventSource("/api/events")


events.onmessage = (event) => {
    console.log("SSE event:", event.data)

    const data = JSON.parse(event.data)

    updateOverview(data)

    data.backends.forEach(backendData => {
        updateBackendHealth(backendData)
        updateBackendMetrics(backendData)
    })
}
