import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Login as LoginAPI, Verify } from "../api/Auth";
import { Credential } from "../types/Credentials";
import "../styles/login.css"

const Login: React.FC = () => {

    const navigate = useNavigate();
    const [credentials, setCredentials] = useState<{ Username: string, Password: string }>({ Username: "", Password: "" });
    const [popoverText, setPopoverText] = useState<string>("");

    useEffect(() => {
        const verify = async () => {
            try {
                const verified = await Verify();
                if (verified) {
                    navigate("/admin");
                }
            } catch (error) {
                console.error(error);
            }
        };
        verify();
    }, []);

    useEffect(() => {
        let timer: number;
        if (popoverText) {
            timer = window.setTimeout(() => {
                setPopoverText("");
            }, 5000);
        }
        return () => window.clearTimeout(timer);

    }, [popoverText]);

    const handleLoginButton = async (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
        const credential: Credential = { Username: credentials.Username, Password: credentials.Password };
        const loggedIn = await LoginAPI(credential);
        if (!loggedIn) {
            setPopoverText("There was an issue logging in.")
            return;
        }

        setCredentials({ Username: "", Password: "" });
        navigate("/admin");
    }

    const handleCredentialChange = (field: string, value: string) => {
        setCredentials((prevState) => ({
            ...prevState,
            [field]: value
        }));
    }

    const enableLoginButton = () => {
        const { Username, Password } = credentials || {};
        const isUsernameValid = /^[a-zA-Z\d]+$/.test(Username);
        const isPasswordValid = /^[a-zA-Z\d]+$/.test(Password);
        return isUsernameValid && isPasswordValid;
    }

    return (
        <div className="login-container">
            <main>
                <div style={{ display: popoverText ? 'block' : 'none' }} className="popover">
                    {popoverText}
                </div>
                <form className="login-form">
                    <h1>Coach Login</h1>
                    <input
                        type="text"
                        placeholder="username"
                        value={credentials.Username}
                        onChange={(e) => handleCredentialChange("Username", e.target.value)}
                    />
                    <input
                        type="password"
                        placeholder="password"
                        value={credentials.Password}
                        onChange={(e) => handleCredentialChange("Password", e.target.value)}
                    />
                    <button type="submit" disabled={!enableLoginButton()} onClick={handleLoginButton}>login</button>
                </form>
            </main>
        </div >
    );
};

export default Login;

