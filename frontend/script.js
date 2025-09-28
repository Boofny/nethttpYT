
document.getElementById("But").addEventListener("click", async () => {
	try{
		const resp = await fetch("http://localhost:8081/test")
		const data = await resp.json()
  	document.getElementById("display").textContent = data.hello
    console.log("Good")
	}catch(err){
  	document.getElementById("display").textContent = "Error"
    console.log("Error")
  }
})
