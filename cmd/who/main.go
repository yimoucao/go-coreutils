package main

import "fmt"

/*
NAME
       who - show who is logged on

SYNOPSIS
       who [OPTION]... [ FILE | ARG1 ARG2 ]

DESCRIPTION
       Print information about users who are currently logged in.

       -a, --all
              same as -b -d --login -p -r -t -T -u

       -b, --boot
              time of last system boot

       -d, --dead
              print dead processes

       -H, --heading
              print line of column headings

       --ips  print  ips  instead of hostnames. with --lookup, canonicalizes based on stored IP, if available, rather
              than stored hostname

       -l, --login
              print system login processes
*/

/*
struct utmpx {
	short		ut_type;
	struct timeval	ut_tv;
	char		ut_id[8];
	pid_t		ut_pid;
	char		ut_user[32];
	char		ut_line[16];
#if __BSD_VISIBLE
	char		ut_host[128];
#else
	char		__ut_host[128];
#endif
	char		__ut_spare[64];
};
*/
// https://github.com/freebsd/freebsd/blob/a3bbb67b61eb7fd57b00e602672944da810e7417/include/utmpx.h#L43

func main() {
	who()
}

func who() {
	fmt.Println("not implemented yet")
}
