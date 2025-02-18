import { useEffect, useState } from "react";
import { axiosAPI } from "./utils/axiosAPI";

export default function App() {
    const [UUID, setUUID] = useState<string>("");
    
    const genUUID = async () => {
        const res = await axiosAPI.get("/api/hello");
        setUUID(res.data);
    }

    useEffect(() => {
        genUUID();
    }, []);

    return (
        <div className="flex flex-col items-center justify-center h-screen gap-5">
            <div className="text-3xl">
                Go Webapp Template
            </div>
            <div>{UUID ?? "Loading..."}</div>
            <button type="button" onClick={genUUID} className="px-3 py-2 bg-sky-600 text-white rounded-2xl cursor-pointer hover:bg-sky-700 transition-all">Click me</button>
        </div>
    )
}