'use client'

import {useAuth} from "@/context/auth-context";
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import {Separator} from "@/components/ui/separator";
import {Button} from "@/components/ui/button";
import {LogOut} from "lucide-react";
import {CameraView} from "@/components/camera-view";
import {CreateRoom} from "@/components/create-room";
import {HorizontalOrSeparator} from "@/components/horizontal-or-seperator";
import {JoinRoom} from "@/components/join-room";

export default function Home() {
    const { user, logout } = useAuth()
    return (
    <div className="flex min-h-screen flex-col items-center justify-center ">
        <Card className="w-[900px] h-[500px]">
            <CardHeader>
                <CardTitle>Welcome {user?.name}</CardTitle>
                <CardDescription>Join a room or create a one.</CardDescription>
            </CardHeader>
            <CardContent className="w-full flex h-[300px] items-center space-x-4 text-sm">
                <div className="w-1/2 flex flex-col gap-4">
                    <CreateRoom />
                    <HorizontalOrSeparator />
                    <JoinRoom />
                </div>
                <Separator orientation="vertical"  />
                <div className="w-1/2 block h-full">
                    <CameraView />
                </div>
            </CardContent>
            <CardFooter className="flex items-end justify-end gap-4">
                <Button variant="ghost" className="cursor-pointer text-red-400 items-end" onClick={logout}>
                    <span>Sign out</span>
                    <LogOut className="size-4" />
                </Button>
            </CardFooter>
        </Card>
    </div>
  );
}
