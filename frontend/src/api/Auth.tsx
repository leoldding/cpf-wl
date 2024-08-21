import { Credential } from "../types/Credentials"

const API_URL = import.meta.env.VITE_WL_API_URL

export async function Verify(): Promise<boolean> {
    try {
        const response = await fetch(API_URL + "/verify", {
            method: "GET",
            credentials: "include",
        });
        if (!response.ok) {
            throw new Error("Unable to verify token");
        }
        return true;
    } catch (error) {
        return false;
    }
}

export async function Login(credential: Credential): Promise<boolean> {
    try {
        const response = await fetch(API_URL + "/login", {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(credential),
        });
        if (!response.ok) {
            throw new Error("Unable to login");
        }
        return true;
    } catch (error) {
        return false;
    }
}

export async function Logout(): Promise<boolean> {
    try {
        const response = await fetch(API_URL  + "/logout", {
            method: "GET",
            credentials: "include",

        });
        if (!response.ok) {
            throw new Error("Unable to logout");
        }
        return true;
    } catch (error) {
        return false;
    }
}
