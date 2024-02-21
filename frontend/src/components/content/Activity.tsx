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
  const [hoverStates, setHoverStates] = useState<Record<number, boolean>>({});

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

          // show pulseRing00 if activity is NOT stale, else show Statle + Remove button
          const stale = (id: number) => {
            if (r.stale_flag) {
              return (
                <div
                  id="staleButton"
                  key={id}
                  onMouseEnter={() => {
                    setHoverStates(prevHoverStates => ({
                      ...prevHoverStates,
                      [id]: true
                    }));
                    console.log(`Button ${id} is being hovered over.`);
                  }}
                  onMouseLeave={() => {
                    setHoverStates(prevHoverStates => ({
                      ...prevHoverStates,
                      [id]: false
                    }));
                    console.log(`Button ${id} is no longer being hovered over.`);
                  }}
                  onClick={() => {
                    console.log(`Button ${r.url} was clicked.`);
                    App.RemoveStaleActivity(r.url)
                  }}
                >
                  {hoverStates[id] ? "Remove" : "Stale"}
                </div>
              );
            } else {
              return (
                <img className="activityLoadingSVG" src={pulseRing00}></img>
              );
            }
          };

          return(
            <div className="activityRow" key={i}>
              <span className="activityRowGroup">
                <img src={downloadingSVG} alt="Downloading"></img>
                {favicon}
                {r.url}
              </span>
              <span className="activityRowGroup">
                {stale(i)}
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
    }, 1000000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div id="activity">
      {jobRunning ? <>{activity}</> : <p>Job NOT Running</p>}
    </div>
  );
}

export default Activity;
