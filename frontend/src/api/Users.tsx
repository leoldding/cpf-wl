import { User } from "../types/User"

export async function CreateUser(user: User): Promise<User> {
    try {
        const response = await fetch("http://localhost:8080/api/users", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(user),
        });
        if (!response.ok) {
            throw new Error(`Error creating user $user.name`);
        }
        const returnedUser: User = await response.json();
        return returnedUser;
    } catch (error) {
        console.error(error);
        throw(error);
    }
}

export async function GetUsers(): Promise<User[]> {
    try {
        const response = await fetch("http://localhost:8080/api/users", {
            method: "GET",
        });
        if (!response.ok) {
            throw new Error("Error retrieving users");
        }
        const returnedUsers: User[] = await response.json();
        return returnedUsers;
    } catch (error) {
        console.error(error);
        throw(error);
    }
}

export async function UpdateUser(user: User): Promise<boolean> {
    try {
        const response = await fetch("http://localhost:8080/api/users", {
            method: "PATCH",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(user),
        });
        if (!response.ok) {
            throw new Error(`Error updating using $user.name`);
        }
        return true;
    } catch (error) {
        console.error(error);
        return false;
    }
}

export async function DeleteUser(user: User): Promise<boolean> {
    try {
        const response = await fetch("http://localhost:8080/api/users/" + user.Id, {
            method: "DELETE",
        });
        if (!response.ok) {
            throw new Error(`Error deleting user $user.name`);
        }
        return true;
    } catch (error) {
        console.error(error);
        return false;
    }
}
