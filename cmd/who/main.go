package main

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
