import "./Settings.css";
import * as App from "../../../wailsjs/go/main/App";
import { useEffect, useRef, useState } from "react";

function Settings() {

  const [contentDir, setContentDir] = useState("");

  const inputRef = useRef("");

  // get contetn dir from settings and display
  App.GetSettings().then((data) => {
    // set returned values
    setContentDir(data.content_dir);
  });

  const handleClick = () => {
    // update the content dir based on user selection
    App.UpdateContentDir().then(() => {
      // update content dir in frontend
      App.GetSettings().then((data) => {
        setContentDir(data.content_dir);
      });
    })
  };
  const handleBlur = (e: React.FocusEvent<HTMLInputElement, Element>) => {
    e.target.placeholder = "Enter URL";
  };
  const handleFocus = (e: React.FocusEvent<HTMLInputElement, Element>) => {
    e.target.placeholder = "";
  };
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    inputRef.current = e.target.value;
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
