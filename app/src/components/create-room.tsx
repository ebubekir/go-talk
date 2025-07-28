import {Button} from "@/components/ui/button";
import {useState} from "react";
import {showErrorToastMessage} from "@/components/ui/toast";
import {LoaderIcon} from "lucide-react";
import {api} from "@/api/api";
import {useRouter} from "next/navigation";

export function CreateRoom() {
    const [isLoading, setIsLoading] = useState(false)
    const router = useRouter()

    const onCreateRoomClick = async () => {
        try {
            setIsLoading(true)
            const response = await api.rooms.create()
            router.push(`/room/${response.id}`)
        } catch (e) {
            showErrorToastMessage("An error occurred while creating the room.")
        }
    }

    return (
        <Button variant="outline" className="w-full h-16 cursor-pointer" disabled={isLoading} onClick={onCreateRoomClick}>
            {isLoading && <LoaderIcon className="animate-spin" />}
            {isLoading ? "Please wait..." : "Create Room"}
        </Button>
    )
}