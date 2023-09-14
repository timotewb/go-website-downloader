import { lib } from "../../../../wailsjs/go/models";
import "./shared.css";
import successSVG from "../../../assets/images/success.svg";

function Found(props: lib.ResponseType) {
  return (
    <>
      <div id="messageArea">
        <div className="icon">
          <img src={successSVG}></img>
        </div>
        <div className="message">Found!</div>
      </div>
      <div className="favicon">
        <img src={props.favicon_url}></img>
        <p>{props.url}</p>
      </div>
    </>
  );
}

export default Found;
