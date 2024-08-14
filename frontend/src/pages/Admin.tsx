import React, { useEffect, useState } from "react";
import { User } from "../types/User"
import { GetUsers, UpdateUser, DeleteUser } from "../api/Users"
import "../styles/admin-leaderboard.css"
import { IoMdCheckmark } from "react-icons/io";
import { FaRegTrashAlt } from "react-icons/fa";

const Admin: React.FC = () => {

    const [users, setUsers] = useState<User[]>([]);
    const [inputValues, setInputValues] = useState<{ [key: string]: { Name: string, Snatch: string, CleanJerk: string } }>({});

    const sortUsers = (users: User[]) => {
        return users.sort((a, b) => a.Name.localeCompare(b.Name))
    };

    useEffect(() => {
        const getUsers = async () => {
            try {
                const users = await GetUsers()
                setUsers(sortUsers(users));

                const initialValues = users.reduce((acc, user) => {
                    acc[user.Id] = { Name: "", Snatch: "", CleanJerk: "" };
                    return acc;
                }, {} as { [key: string]: { Name: string, Snatch: string, CleanJerk: string } });
                setInputValues(initialValues);
            } catch (error) {
                console.error(error);
            }
        };
        getUsers();
    }, []);

    const handleDeleteButton = (userId: string) => {
        DeleteUser(userId);

        const newUsers = users.filter(user => user.Id != userId);
        setUsers(newUsers);
    }

    const handleUpdateButton = (userId: string) => {
        const name = inputValues[userId].Name;
        const snatch = inputValues[userId].Snatch;
        const cleanjerk = inputValues[userId].CleanJerk;

        let user: User = {} as User;
        const newUsers = [...users];
        let index: number = 0;
        for (let i = 0; i < newUsers.length; i++) {
            if (newUsers[i].Id === userId) {
                user = newUsers[i];
                index = i;
                break;
            }
        }
        if (name !== "") {
            user.Name = name;
            inputValues[userId].Name = "";
        }
        if (snatch !== "") {
            user.Snatch = parseInt(snatch);
            inputValues[userId].Snatch = "";
        }
        if (cleanjerk !== "") {
            user.CleanJerk = parseInt(cleanjerk);
            inputValues[userId].CleanJerk = "";
        }
        if (snatch !== "" || cleanjerk !== "") {
            user.Total = user.Snatch + user.CleanJerk;
        }
        UpdateUser(user);

        newUsers[index] = user;
        setUsers(sortUsers(newUsers));
    }

    const handleInputChange = (userId: string, field: keyof User, value: string) => {
        setInputValues((prevState) => ({
            ...prevState,
            [userId]: {
                ...prevState[userId],
                [field]: value
            }
        }));
    }

    const enableButton = (userId: string) => {
        const { Name, Snatch, CleanJerk } = inputValues[userId] || {};
        const isNameValid = Name === "" || /^[a-zA-Z\s]+$/.test(Name) && !/^\s+$/.test(Name);
        const isSnatchValid = Snatch === "" || /^\d+$/.test(Snatch);
        const isCleanJerkValid = CleanJerk === "" || /^\d+$/.test(CleanJerk);

        return (!!Name || !!Snatch || !!CleanJerk) && isNameValid && isSnatchValid && isCleanJerkValid;
    }

    return (
        <div className="admin-container">
            <main>
                <div className="admin-leaderboard">
                    <div className="admin-leaderboard-header">
                        <div>Name</div>
                        <div>Snatch</div>
                        <div>Clean & Jerk</div>
                        <div>Total</div>
                        <div />
                    </div>
                    {users.map(user => (<div key={user.Id} className="admin-leaderboard-row">
                        <input
                            type="text"
                            placeholder={user.Name}
                            value={inputValues[user.Id]?.Name || ""}
                            onChange={(e) => handleInputChange(user.Id, "Name", e.target.value)}
                        />
                        <input
                            type="text"
                            placeholder={user.Snatch.toString()}
                            value={inputValues[user.Id]?.Snatch || ""}
                            onChange={(e) => handleInputChange(user.Id, "Snatch", e.target.value)}
                        />
                        <input
                            type="text"
                            placeholder={user.CleanJerk.toString()}
                            value={inputValues[user.Id]?.CleanJerk || ""}
                            onChange={(e) => handleInputChange(user.Id, "CleanJerk", e.target.value)}
                        />
                        <div>{user.Total}</div>
                        <div>
                            <button type="button" disabled={!enableButton(user.Id)} onClick={() => handleUpdateButton(user.Id)}><IoMdCheckmark /></button>
                            <button type="button" onClick={() => handleDeleteButton(user.Id)}><FaRegTrashAlt /></button>
                        </div>
                    </div>))}
                </div>
            </main>
        </div>
    );
};

export default Admin;
