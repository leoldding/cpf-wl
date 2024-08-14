import React, { useEffect, useState } from "react";
import { User } from "../types/User";
import { GetUsers } from "../api/Users";
import "../styles/leaderboard.css";

const Main: React.FC = () => {

    const [users, setUsers] = useState<User[]>([]);

    useEffect(() => {
        const getUsers = async () => {
            try {
                const users = await GetUsers()
                setUsers(users);
            } catch (error) {
                console.error(error);
            }
        };
        getUsers();
    }, []);

    return (
        <div className="main-container">
            <main>
                <div className="leaderboard">
                    <div className="leaderboard-header">
                        <div>Name</div>
                        <div>Snatch</div>
                        <div>Clean & Jerk</div>
                    </div>
                    {users.map(user => (<div key={user.Id} className="leaderboard-row">
                        <div>{user.Name}</div>
                        <div>{user.Snatch}</div>
                        <div>{user.CleanJerk}</div>
                    </div>))}
                </div>
            </main>
        </div>
    );
};

export default Main;
