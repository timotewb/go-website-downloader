import "./Settings.css";
import * as App from "../../../wailsjs/go/main/App";
import { useEffect, useRef, useState } from "react";

function Settings() {

  const [contentDir, setContentDir] = useState("");

  const inputRef = useRef("");

  App.GetSettings().then((data) => {
    // set returned values
    setContentDir(data.content_dir);
  });

  const handleClick = () => {
    console.log("---- testing ----");
  };
  return <>
        <div id="input">
        <input
          disabled={true}
          id="inputArea"
          placeholder={contentDir}
        ></input>
        <div id="inputButton" onClick={() => handleClick()}>
          Update
        </div>
      </div>
  </>;
}

export default Settings;
