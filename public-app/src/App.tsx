import { useEffect } from "react";
import "./App.css";
import { AppState, appStore } from "./core/store";

function App() {
  const fetchData = appStore(
    (state: AppState) => state.accommodation.fetchAccommodations
  );
  const data = appStore((state: AppState) => state.accommodation.data);

  useEffect(() => {
    fetchData();
  }, [fetchData]);
  return (
    <div className="App">
      AAAAAAAAA
      {data.map((d, i) => (
        <div key={i}>{d.name}</div>
      ))}
    </div>
  );
}

export default App;
