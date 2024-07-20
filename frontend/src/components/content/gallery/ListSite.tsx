import { PageContext, PageContextType } from "../../../App";
import { useContext, useState, useEffect } from "react";
import "../shared.css";
import "../Activity.css";
import * as App from "../../../../wailsjs/go/main/App";
import { models } from "../../../../wailsjs/go/models";

interface ListSiteProps {
  site_name: string;
}

function formatDate(dateTimeString: string): string {
  // Extract year, month, day, hours, minutes, and AM/PM indicator
  const [datePart, timePart] = dateTimeString.split("_");
  const year = datePart.substring(0, 4);
  const month = datePart.substring(4, 6);
  const day = datePart.substring(6);
  const hours = Number(timePart.substring(0, 2));
  const minutes = Number(timePart.substring(2, 4));
  const ampm = hours >= 12 ? "pm" : "am";

  // Convert hours to 12-hour format and adjust for AM/PM
  const hour = hours % 12 || 12;

  // Pad month and day with leading zeros if necessary
  const formattedMin = String(minutes).padStart(2, "0");
  const formattedHour = String(hours).padStart(2, "0");

  // Construct the final date string
  return `${year}-${month}-${day} ${formattedHour}:${formattedMin}${ampm}`;
}

function ListSite(props: ListSiteProps) {
  const pageContext: PageContextType = useContext(PageContext);
  const [gallerySite, setGallerySite] = useState<JSX.Element[] | null>(null);

  const handleClick = () => {
    pageContext.gallery.setShowSiteList(false);
  };
  useEffect(() => {
    App.ListGallerySite(props.site_name).then((data) => {
      console.log("App.ListGallerySite()");

      data.sort((a: models.GallerySiteType, b: models.GallerySiteType) =>
        b.date_time.localeCompare(a.date_time)
      );

      const gallerySite = data.map((s: models.GallerySiteType, i: number) => {
        return (
          <div className="activityRow" key={i}>
            <span
              className="listSiteRowGroup"
              onClick={() => {
                App.OpenSite(props.site_name, s.date_time);
              }}
            >
              <img id="faviconImgRow" src={s.favicon}></img>
              {formatDate(s.date_time)}
            </span>
          </div>
        );
      });
      setGallerySite(gallerySite);
    });
  }, []);

  return (
    <>
      <div id="input">
        <div id="inputButton" onClick={handleClick}>
          &lt;&lt;
        </div>
      </div>
      {props.site_name}
      <div id="activity">{gallerySite}</div>
    </>
  );
}

export default ListSite;
