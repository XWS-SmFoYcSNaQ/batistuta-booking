import axios from "axios";
import { useState, useEffect, useRef } from "react";

interface State {
  loading: boolean;
  error: string | null;
  data: any;
}

function useFetch (url: string) {
  const [state, setState] = useState<State>({ loading: false, error: null, data: null});

  const axiosConfig = useRef({
    headers: {
      "Content-Type" : "application/json",
      "Authorization" : `Bearer ${window.localStorage.getItem('jwt')}`
    }
  });

  useEffect(() => {
    const fetchData = async () => {
      setState({ loading: true, error: null, data: null});
      try {
        const res = await axios.get(url, axiosConfig.current);
        if (res.status !== 200) {
          console.log((res && res.data.message) ? res.data.message : "Error fetching data.");
          setState(state => {return { ...state, loading: false, error: "Error fetching data." }});
        }
        setState(state => { return { ...state, loading: false, data: res.data }});
      } catch (err: any) {
        console.log(err);
        setState(state => { return { ...state, loading: false, error: "Error fetching data." }});
      }
    }
    fetchData();

  }, [url]);

  return state;
}

export default useFetch;