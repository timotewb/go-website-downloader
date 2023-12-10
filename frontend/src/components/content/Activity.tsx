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
import pulseRing01 from "../../assets/images/pulse-ring-01.svg";
import pulseRing02 from "../../assets/images/pulse-ring-02.svg";
import pulseRing03 from "../../assets/images/pulse-ring-03.svg";

function getRandomElement<T>(arr: T[]): T | undefined {
  return arr[Math.floor(Math.random() * arr.length)];
}

function Activity() {
  const [jobRunning, setJobRunning] = useState(false);
  const [activity, setActivity] = useState<JSX.Element[] | null>(null);
  const pulseRings = [pulseRing00, pulseRing01, pulseRing02, pulseRing03];

  const check = () => {
    App.CheckActivity().then((data) => {
      if (data.job_count > 0) {
        const activity = data.data.map((r, i) => (
          <div className="activityRow">
            <span className="activityRowGroup">
              <img src={downloadingSVG}></img>
              {r.favicon_url ? (
                <img className="activityFavicon" src={r.favicon_url}></img>
              ) : (
                <img className="activityFavicon" src={faviconDefaultSVG}></img>
              )}
              {r.url}
            </span>
            <span className="activityRowGroup">
              <img className="activityLoadingSVG" src={pulseRing00}></img>
            </span>
          </div>
        ));
        setActivity(activity);
        setJobRunning(true);
      } else {
        setJobRunning(false);
      }
    });
  };

  // run for te first time the page loads
  check();

  // run every 1 second
  useEffect(() => {
    const interval = setInterval(() => {
      check();
    }, 10000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div id="activity">
      {jobRunning ? <>{activity}</> : <p>Job NOT Running</p>}
    </div>
  );
}

export default Activity;
