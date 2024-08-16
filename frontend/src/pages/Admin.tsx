import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { User } from "../types/User";
import { CreateUser, GetUsers, UpdateUser, DeleteUser } from "../api/Users";
import { Verify, Logout } from "../api/Auth";
import "../styles/admin-leaderboard.css";
import { IoMdCheckmark, IoMdAdd } from "react-icons/io";
import { FaRegTrashAlt } from "react-icons/fa";

const Admin: React.FC = () => {

    const [users, setUsers] = useState<User[]>([]);
    const [inputValues, setInputValues] = useState<{ [key: string]: { Name: string, Snatch: string, CleanJerk: string } }>({});
    const [addUser, setAddUser] = useState<{ Name: string, Snatch: string, CleanJerk: string }>({ Name: "", Snatch: "", CleanJerk: "" });
    const navigate = useNavigate();

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

    useEffect(() => {
        const verify = async () => {
            try {
                const verified = await Verify();
                if (!verified) {
                    navigate("/login");
                }
            } catch (error) {
                console.error(error);
            }
        };
        verify();
    }, []);

    const handleUpdateButton = async (userId: string) => {
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
        const updated = await UpdateUser(user);
        if (!updated) {
            return;
        }

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

    const enableUpdateButton = (userId: string) => {
        const { Name, Snatch, CleanJerk } = inputValues[userId] || {};
        const isNameValid = Name === "" || /^[a-zA-Z\s]+$/.test(Name) && !/^\s+$/.test(Name);
        const isSnatchValid = Snatch === "" || /^\d+$/.test(Snatch);
        const isCleanJerkValid = CleanJerk === "" || /^\d+$/.test(CleanJerk);

        return (!!Name || !!Snatch || !!CleanJerk) && isNameValid && isSnatchValid && isCleanJerkValid;
    }

    const handleDeleteButton = async (userId: string) => {
        const deleted = await DeleteUser(userId);
        if (!deleted) {
            return;
        }
        const newUsers = users.filter(user => user.Id != userId);
        setUsers(newUsers);
    }

    const handleAddButton = async () => {
        const user: User = { Id: "", Total: 0, Name: addUser.Name, Snatch: parseInt(addUser.Snatch), CleanJerk: parseInt(addUser.CleanJerk) }
        const newUser = await CreateUser(user);
        if (!newUser) {
            return;
        }

        setAddUser({ Name: "", Snatch: "", CleanJerk: "" });
        const newUsers = [...users];
        newUsers.push(newUser);
        setUsers(newUsers);
    }

    const handleAddChange = (field: string, value: string) => {
        setAddUser((prevState) => ({
            ...prevState,
            [field]: value
        }));
    }

    const enableAddButton = () => {
        const { Name, Snatch, CleanJerk } = addUser || {};
        const isNameValid = /^[a-zA-Z\s]+$/.test(Name) && !/^\s+$/.test(Name);
        const isSnatchValid = /^\d+$/.test(Snatch);
        const isCleanJerkValid = /^\d+$/.test(CleanJerk);

        return isNameValid && isSnatchValid && isCleanJerkValid;
    }

    const handleLogout = async () => {
        const loggedOut = await Logout();
        if (loggedOut) {
            navigate("/login");
        }
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
                            <button type="button" disabled={!enableUpdateButton(user.Id)} onClick={() => handleUpdateButton(user.Id)}><IoMdCheckmark /></button>
                            <button type="button" onClick={() => handleDeleteButton(user.Id)}><FaRegTrashAlt /></button>
                        </div>
                    </div>))}
                    <div className="admin-leaderboard-row">
                        <input
                            type="text"
                            placeholder="New Athlete"
                            value={addUser.Name}
                            onChange={(e) => handleAddChange("Name", e.target.value)}
                        />
                        <input
                            type="text"
                            placeholder="Snatch Weight"
                            value={addUser.Snatch}
                            onChange={(e) => handleAddChange("Snatch", e.target.value)}
                        />
                        <input
                            type="text"
                            placeholder="Clean & Jerk Weight"
                            value={addUser.CleanJerk}
                            onChange={(e) => handleAddChange("CleanJerk", e.target.value)}
                        />
                        <div />
                        <div>
                            <button type="button" disabled={!enableAddButton()} onClick={handleAddButton}><IoMdAdd /></button>
                        </div>
                    </div>
                </div>
            </main>
            <button type="button" onClick={handleLogout}>logout</button>
        </div>
    );
};

export default Admin;
