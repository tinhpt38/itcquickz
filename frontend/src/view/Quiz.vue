<template>
    <div class="grid grid-cols-5 gap-4">
        <div class="bg-white drop-shadow-xl">
            <section class="relative isolate overflow-hidden bg-white px-6 py-24 sm:py-32 lg:px-8">
                <div
                    class="absolute inset-0 -z-10 bg-[radial-gradient(45rem_50rem_at_top,theme(colors.indigo.100),white)] opacity-20" />
                <div
                    class="absolute inset-y-0 right-1/2 -z-10 mr-16 w-[200%] origin-bottom-left skew-x-[-30deg] bg-white shadow-xl shadow-indigo-600/10 ring-1 ring-indigo-50 sm:mr-28 lg:mr-0 xl:mr-16 xl:origin-center" />
                <div class="text-8xl text-center mt-2 mb-2">
					<span>{{ timer.minutes }}</span
                    >:<span>{{ timer.seconds }}</span>
                </div>
                <div class="mx-auto max-w-2xl lg:max-w-4xl">
                    <figure class="mt-10">
                        <figcaption class="mt-10">
                            <!--<img-->
                            <!--  v-if="user.sex == 'Nam'"-->
                            <!--  class="mx-auto h-10 w-10 rounded-full"-->
                            <!--  src="./../assets/images/man.png"-->
                            <!--  alt=""-->
                            <!--/>-->
                            <img class="mx-auto h-10 w-10 rounded-full" src="./../assets/images/woman.png" alt="" />
                            <info-line title="Số báo danh" subTitle="user.idNumber || ''"></info-line>
                            <info-line title="Họ và tên" subTitle="user.fullName || ''"></info-line>
                            <info-line title="Giới tính" subTitle="user.sex || ''"></info-line>
                            <info-line title="Ngày sinh" subTitle="user.bod || ''"></info-line>
                            <info-line title="Nơi sinh" subTitle="user.pod || ''"></info-line>
                            <info-line title="Phòng thi" subTitle="user.room || ''"></info-line>
                        </figcaption>
                    </figure>
                </div>
            </section>
        </div>
        <div class="bg-white col-span-3">
            <ul role="list" class="divide-y divide-gray-100">
                <li v-for="person in people" :key="person.email" class="flex justify-between gap-x-6 py-5">
                    <div class="flex min-w-0 gap-x-4">
                        <img class="h-12 w-12 flex-none rounded-full bg-gray-50" :src="person.imageUrl" alt="" />
                        <div class="min-w-0 flex-auto">
                            <p class="text-sm font-semibold leading-6 text-gray-900">
                                {{ person.name }}
                            </p>
                            <p class="mt-1 truncate text-xs leading-5 text-gray-500">
                                {{ person.email }}
                            </p>
                        </div>
                    </div>
                    <div class="hidden shrink-0 sm:flex sm:flex-col sm:items-end">
                        <p class="text-sm leading-6 text-gray-900">{{ person.role }}</p>
                        <p v-if="person.lastSeen" class="mt-1 text-xs leading-5 text-gray-500">
                            Last seen
                            <time :datetime="person.lastSeenDateTime">{{ person.lastSeen }}</time>
                        </p>
                        <div v-else class="mt-1 flex items-center gap-x-1.5">
                            <div class="flex-none rounded-full bg-emerald-500/20 p-1">
                                <div class="h-1.5 w-1.5 rounded-full bg-emerald-500" />
                            </div>
                            <p class="text-xs leading-5 text-gray-500">Online</p>
                        </div>
                    </div>
                </li>
            </ul>
        </div>
        <div class="bg-white drop-shadow-xl">
            <section class="relative isolate overflow-hidden bg-white px-6 py-24 sm:py-32 lg:px-8">
                <div
                    class="absolute inset-0 -z-10 bg-[radial-gradient(45rem_50rem_at_top,theme(colors.indigo.100),white)] opacity-20"
                />
                <div
                    class="absolute inset-y-0 right-1/2 -z-10 mr-16 w-[200%] origin-bottom-left skew-x-[-30deg] bg-white shadow-xl shadow-indigo-600/10 ring-1 ring-indigo-50 sm:mr-28 lg:mr-0 xl:mr-16 xl:origin-center"
                />
            </section>
        </div>
        <app-dialog
            ref="dialogRef"
            title="Hết thời gian làm bài"
            detail=""
            :showAction="false"
            @onDismissClick=""
            @onOpen=""
        ></app-dialog>
    </div>
</template>

<script setup>
import InfoLine from "../components/InfoLine.vue";
import { ref } from "vue";
import { watchEffect, onMounted } from "vue";
import { useTimer } from "vue-timer-hook";
import AppDialog from "../components/AppDialog.vue";

const dialogRef = ref();
const time = new Date();
time.setSeconds(time.getSeconds() + 60); // 1 minutes timer
const timer = useTimer(time);
//const restartFive = () => {
//  // Restarts to 5 minutes timer
//  const time = new Date();
//  time.setSeconds(time.getSeconds() + 300);
//  timer.restart(time);
//};
function onAutoSubmit() {
    //  TODO: onSumit
}

onMounted(() => {
    watchEffect(async () => {
        if (timer.isExpired.value) {
            console.log("IsExpired");
            dialogRef.value.showDialog();
        }
    });
});

const people = [
    {
        name: "Leslie Alexander",
        email: "leslie.alexander@example.com",
        role: "Co-Founder / CEO",
        imageUrl:
            "https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
        lastSeen: "3h ago",
        lastSeenDateTime: "2023-01-23T13:23Z"
    },
    {
        name: "Michael Foster",
        email: "michael.foster@example.com",
        role: "Co-Founder / CTO",
        imageUrl:
            "https://images.unsplash.com/photo-1519244703995-f4e0f30006d5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
        lastSeen: "3h ago",
        lastSeenDateTime: "2023-01-23T13:23Z"
    },
    {
        name: "Dries Vincent",
        email: "dries.vincent@example.com",
        role: "Business Relations",
        imageUrl:
            "https://images.unsplash.com/photo-1506794778202-cad84cf45f1d?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
        lastSeen: null
    },
    {
        name: "Lindsay Walton",
        email: "lindsay.walton@example.com",
        role: "Front-end Developer",
        imageUrl:
            "https://images.unsplash.com/photo-1517841905240-472988babdf9?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
        lastSeen: "3h ago",
        lastSeenDateTime: "2023-01-23T13:23Z"
    },
    {
        name: "Courtney Henry",
        email: "courtney.henry@example.com",
        role: "Designer",
        imageUrl:
            "https://images.unsplash.com/photo-1438761681033-6461ffad8d80?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
        lastSeen: "3h ago",
        lastSeenDateTime: "2023-01-23T13:23Z"
    },
    {
        name: "Tom Cook",
        email: "tom.cook@example.com",
        role: "Director of Product",
        imageUrl:
            "https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
        lastSeen: null
    }
];
</script>
