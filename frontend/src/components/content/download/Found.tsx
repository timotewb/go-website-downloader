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
      <div className="faviconDownloadArea">
        <img id="faviconImg" src={props.favicon_url}></img>
        <div id="downloadNowButton">
          <div id="inputButton">Download</div>
        </div>
      </div>
    </>
  );
}

export default Found;
