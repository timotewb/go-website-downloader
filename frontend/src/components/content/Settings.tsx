import "./Settings.css";
import * as App from "../../../wailsjs/go/main/App";
import { useEffect, useRef, useState, useCallback } from "react";

function Settings() {
  const [contentDir, setContentDir] = useState("");
  const [wsPort, setWsPort] = useState("");
  const [wsPortIsValid, setWsPortIsValid] = useState(false);

  // get contetn dir from settings and display
  const getSettings = () => {
    App.GetSettings().then((data) => {
      // set returned values
      setContentDir(data.content_dir);
      setWsPort(data.content_dir_wsport.toString());
    });
  };

  // run etSettings once
  useEffect(() => {
    getSettings();
  }, []);

  const handleClickCD = () => {
    // update the content dir based on user selection
    App.UpdateContentDir().then(() => {
      // update content dir in frontend
      App.GetSettings().then((data) => {
        setContentDir(data.content_dir);
      });
    });
  };

  const handleClickWSP = () => {
    if (wsPortIsValid) {
      console.log("wsPort:", wsPort);
      App.UpdatePortNumber(parseInt(wsPort));
    }
  };
  const handleChangeWSP = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const port = e.target.value;
      setWsPort(port.toString());
      setWsPortIsValid(validatePort(port));
    },
    []
  );
  useEffect(() => {
    console.log("useEffect:", wsPort); // This will log the updated value of wsPort
  }, [wsPort]);

  const validatePort = (port: string): boolean => {
    const portNum = parseInt(port, 10);
    return portNum >= 1024 && portNum <= 65535;
  };

  return (
    <>
      <div className="setting">
        Content Directory
        <div id="input">
          <input
            disabled={true}
            id="inputArea"
            placeholder={contentDir}
          ></input>
          <div id="inputButton" onClick={() => handleClickCD()}>
            Update
          </div>
        </div>
      </div>
      <div className="setting">
        Webserver Port
        <div id="input">
          <input
            id="inputArea"
            style={{ maxWidth: "50px" }}
            placeholder={wsPort}
            onChange={handleChangeWSP}
          ></input>

          <div
            id="inputButton"
            style={{ width: "12px" }}
            onClick={() => handleClickWSP()}
          >
            Set
          </div>
        </div>
      </div>
    </>
  );
}

export default Settings;
