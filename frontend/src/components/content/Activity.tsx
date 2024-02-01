import {
  JSXElementConstructor,
  ReactElement,
  ReactNode,
  useEffect,
  useState,
} from "react";
import "./Activity.css";
import "./shared.css";
import * as App from "../../../wailsjs/go/main/App";
import { JSX } from "react/jsx-runtime";
import downloadingSVG from "../../assets/images/downloading.svg";
import faviconDefaultSVG from "../../assets/images/favicon-default.svg";
import pulseRing00 from "../../assets/images/pulse-ring-00.svg";

function Activity() {
  const [jobRunning, setJobRunning] = useState(false);
  const [activity, setActivity] = useState<JSX.Element[] | null>(null);
  const [ isHovered, setIsHovered ] = useState(false);

  const check = () => {
    App.CheckActivity().then((data) => {
      if (data.job_count > 0) {
        const activity = data.data.map((r, i) => {

          // use favicon url, or default
          const favicon = r.favicon_url ? (
            <img className="activityFavicon" src={r.favicon_url}></img>
          ) : (
            <img className="activityFavicon" src={faviconDefaultSVG}></img>
          );

          // show pulseRing00 if activity is NOT stale, else show Statle + Delete button
          const stale = () => {
            if (r.stale_flag){
              return(
                <span onMouseEnter={() => setIsHovered(true)} onMouseLeave={() => setIsHovered(false)}>
                  {isHovered ?         <div
          onClick={() => console.log("sdf")}>
          <div id="staleButton">Remove</div>
        </div> : <div id="staleButton">Stale</div>}
                </span>
              )
            } else {
              return(<img className="activityLoadingSVG" src={pulseRing00}></img>)
            }
          }

          return(
            <div className="activityRow" key={i}>
              <span className="activityRowGroup">
                <img src={downloadingSVG} alt="Downloading"></img>
                {favicon}
                {r.url}
              </span>
              <span className="activityRowGroup">
                {stale()}
              </span>
            </div>
          )
      });
        setActivity(activity);
        setJobRunning(true);
      } else {
        setJobRunning(false);
      }
    });
  };

  // run for the first time the page loads
  check();

  // run every 1 second
  useEffect(() => {
    const interval = setInterval(() => {
      check();
    }, 100000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div id="activity">
      {jobRunning ? <>{activity}</> : <p>Job NOT Running</p>}
    </div>
  );
}

export default Activity;
