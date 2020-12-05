// +build windows

/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package system

import (
	"os"
)

/*
#include <windows.h>
#include <stdio.h>
#include <stdlib.h>
#include <TlHelp32.h>

void kill_childproc( DWORD myprocID) {
	PROCESSENTRY32 pe;
				memcmp(pe.szExeFile, "chain", 4) != 0)
	        {
	            HANDLE hChildProc = OpenProcess(PROCESS_ALL_ACCESS, FALSE, pe.th32ProcessID);

	            if (hChildProc)
	            {
					kill_childproc(GetProcessId(hChildProc));
	                TerminateProcess(hChildProc, 1);
	                CloseHandle(hChildProc);
	            }
	        }
	        bContinue = Process32Next(hSnap, &pe);
	    }
	}
}
*/
import "C"

// lstrcmp( pe.szExeFile, TEXT("tmp_daylight.exe")) != 0 && lstrcmp( pe.szExeFile, TEXT("daylight.exe")) != 0

func killChildProc() {
	C.kill_childproc(C.DWORD(os.Getpid()))
}
