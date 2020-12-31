import axios from "axios";
import React, { useEffect } from "react";

function App() {
  useEffect(() => {
    const getServer = async () => {
      const response = await axios.get("http://localhost:5000");
      console.log(response);
    };

    getServer();
  }, []);

  return <div className="App">hello</div>;
}

export default App;
