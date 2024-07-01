import { PageContext, PageContextType } from "../../../App";
import { useContext, useState, useEffect } from "react";
import "../shared.css";
import "../Activity.css";
import * as App from "../../../../wailsjs/go/main/App";
import { models } from "../../../../wailsjs/go/models";

interface ListSiteProps {
  site_name: string;
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
      const gallerySite = data.map((s: models.GallerySiteType, i: number) => {
        return (
          <div className="activityRow" key={i}>
            <span className="activityRowGroup">
              <img id="faviconImgRow" src={s.favicon}></img>
              {s.date_time}
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
