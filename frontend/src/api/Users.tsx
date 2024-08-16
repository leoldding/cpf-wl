import { User } from "../types/User"

export async function CreateUser(user: User): Promise<User | null> {
    try {
        const response = await fetch("http://localhost:8080/api/users", {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(user),
        });
        if (!response.ok) {
            throw new Error(`Error creating user ${user.Name}`);
        }
        const returnedUser: User = await response.json();
        return returnedUser;
    } catch (error) {
        return null;
    }
}

export async function GetUsers(): Promise<User[] | null> {
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
        return null;
    }
}

export async function UpdateUser(user: User): Promise<boolean> {
    try {
        const response = await fetch("http://localhost:8080/api/users", {
            method: "PATCH",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(user),
        });
        if (!response.ok) {
            throw new Error(`Error updating user ${user.Id}`);
        }
        return true;
    } catch (error) {
        return false;
    }
}

export async function DeleteUser(userId: string): Promise<boolean> {
    try {
        const response = await fetch("http://localhost:8080/api/users/" + userId, {
            method: "DELETE",
            credentials: "include",
        });
        if (!response.ok) {
            throw new Error(`Error deleting user ${userId}`);
        }
        return true;
    } catch (error) {
        return false;
    }
}
