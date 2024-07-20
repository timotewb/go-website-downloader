import { useState, useEffect, useRef } from "react";
import "./Gallery.css";
import "./shared.css";
import * as App from "../../../wailsjs/go/main/App";
import { models } from "../../../wailsjs/go/models";
import { PageContext, PageContextType } from "../../App";
import { useContext } from "react";
import ListSite from "./gallery/ListSite";

function Gallery() {
  const pageContext: PageContextType = useContext(PageContext);
  const [gallery, setGallery] = useState<JSX.Element[] | null>(null);
  const [siteName, setSiteName] = useState("");
  const inputRef = useRef("");

  const getData = (searchTerm: string, sort: number = 0) => {
    const handleClick = (site: string) => {
      setSiteName(site);
      pageContext.gallery.setShowSiteList(true);
    };
    App.ListGallery().then((data) => {
      if (sort ===0){
        // sort data ascending
        data.sort((a: models.GalleryType, b: models.GalleryType) =>
          a.site_name.localeCompare(b.site_name)
        );
      } else if (sort ===1){
        // sort data ascending
        data.sort((a: models.GalleryType, b: models.GalleryType) =>
          b.site_name.localeCompare(a.site_name)
        );
      }
      const gallery = data.map((s: models.GalleryType, i: number) => {
        if (searchTerm.trim().length >= 1) {
          if (s.site_name.includes(searchTerm)) {
            return (
              <div className="galleryCell" key={i}>
                <span
                  onClick={() => handleClick(s.site_name)}
                  className="gallerySiteSelect"
                >
                  <img id="faviconImg" src={s.favicon}></img>
                  <br />
                  {s.site_name}
                </span>
              </div>
            );
          }
        } else {
          return (
            <div className="galleryCell" key={i}>
              <span
                onClick={() => handleClick(s.site_name)}
                className="gallerySiteSelect"
              >
                <img id="faviconImg" src={s.favicon}></img>
                <br />
                {s.site_name}
              </span>
            </div>
          );
        }
      });
      setGallery(gallery);
    });
  };
  const renderContent = () => {
    if (pageContext.gallery.showSiteList) {
      return <ListSite site_name={siteName} />;
    }
    const handleBlur = (e: React.FocusEvent<HTMLInputElement, Element>) => {
      e.target.placeholder = "Search";
    };
    const handleFocus = (e: React.FocusEvent<HTMLInputElement, Element>) => {
      e.target.placeholder = "";
    };
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      getData(e.target.value);
      inputRef.current = e.target.value;
    };
    return (
      <>
        <div id="input">
          <input
            id="inputArea"
            placeholder="Search"
            onFocus={(e) => {
              handleFocus(e);
            }}
            onBlur={(e) => {
              handleBlur(e);
            }}
            onChange={(e) => handleChange(e)}
          ></input>
        </div>
        <div id="gallery">{gallery}</div>
      </>
    );
  };

  useEffect(() => {
    getData("");
    renderContent();
  }, []);

  return renderContent();
}

export default Gallery;
