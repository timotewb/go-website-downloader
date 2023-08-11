import "./App.css";
import MenuBar from "./components/MenuBar";
import Content from "./components/Content";
import { createContext, useState } from "react";

export type PageContextType = {
  download: {
    downloadState: boolean;
    setDownloadState: React.Dispatch<React.SetStateAction<boolean>>;
  };
  gallery: {
    galleryState: boolean;
    setGalleryState: React.Dispatch<React.SetStateAction<boolean>>;
  };
  activity: {
    activityState: boolean;
    setActivityState: React.Dispatch<React.SetStateAction<boolean>>;
  };
  settings: {
    settingsyState: boolean;
    setSettingsState: React.Dispatch<React.SetStateAction<boolean>>;
  };
};
export const PageContext = createContext<PageContextType>(
  {} as PageContextType
);

function App() {
  const [downloadState, setDownloadState] = useState(true);
  const [galleryState, setGalleryState] = useState(false);
  const [activityState, setActivityState] = useState(false);
  const [settingsyState, setSettingsState] = useState(false);

  return (
    <PageContext.Provider
      value={{
        download: { downloadState, setDownloadState },
        gallery: { galleryState, setGalleryState },
        activity: { activityState, setActivityState },
        settings: { settingsyState, setSettingsState },
      }}
    >
      <div id="App">
        <MenuBar />
        <Content />
      </div>
    </PageContext.Provider>
  );
}

export default App;
