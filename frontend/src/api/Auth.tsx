import { Credential } from "../types/Credentials"

export async function Login(credential: Credential): Promise<boolean> {
    try {
        const response = await fetch("http://localhost:8080/api/login", {
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

export async function Verify(): Promise<boolean> {
    try {
        const response = await fetch("http://localhost:8080/api/verify", {
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
