package main

import (
  "os"
  "fmt"
  "path/filepath"
  "io"
  "bufio"
  "strings"
  "regexp"
  "log"
)


/*

  Finds an icon for each input line and prints it

*/

var default_icon = string(0xF054)

func main() {

  reader := bufio.NewReader(os.Stdin)
  // fins ANSI color codes (from the perl module Term:ANSIColor)
  reg, err := regexp.Compile("\033\\[?.*?[\\@-~]")

  if err != nil {
    log.Fatal(err)
  }

  for {
    a, err := reader.ReadString('\n')
    if (err != nil && err == io.EOF) {
      break
    }

    // remove newline
    if len(a) > 0 {
      a = a[:len(a)-1]
    }

    plain := reg.ReplaceAllString(a, "")
    icon := findIcon(plain)
    fmt.Printf("  %s %s\n", icon, a)
  }
}

func findIcon(file string) string {

  if strings.HasPrefix(file, "total") {
    return ""
  }

  // Whole file name
  switch file {
    case "README","LICENSE", "COPYING", "INSTALL", "COPYRIGHT", "AUTHORS", "HISTORY", "CONTRIBUTORS",
         "PATENTS", "VERSION", "NOTICE", "CHANGES": return string(0xE714)
    case "Dockerfile": return string(0xE7B0)
    case "Makefile", "MANIFEST": return string(0xF0AD)
    case "lockfile": return string(0xF023)
    case "pm_to_blib": return string(0xF0AD)
  }

  suffix := filepath.Ext(file)

  if len(suffix) < 1 {
    return default_icon
  }

  // remove dot
  suffix = suffix[1:]

  // Just suffix
  switch suffix {
    case "log", "txt": return string(0xE714)
		case "etx", "info", "markdown", "md", "mkd", "nfo", "pod", "tex", "textile": return string(0xE60E)
    case "json", "msg", "pgn", "rss", "xml", "yml", "RData", "rdata": return string(0xE60B)
		case "pdf", "PDF", "cbr", "cbz", "chm", "djvu": return string(0xF1C1)
    case "odt", "docm", "doc", "docx", "eps", "odb": return string(0xF1C2)
		case "rtf", "odp": return string(0xF035)
		case "pps", "ppt", "pptx", "ppts", "pptxm", "pptsm": return string(0xF1C4)
		case "csv": return string(0xF1C0)
		case "ods", "xla", "xls", "xlsx", "xlsxm", "xltm", "xltx": return string(0xF1C3)
		case "cfg", "conf", "rc", "ini", "viminfo", "pcf", "psf": return string(0xF0AD)
		case "git", "gitignore", "gitattributes", "gitmodules": return string(0xE708)
		case "awk", "bash", "bat", "BAT", "sed", "sh", "zsh", "vim", "ahk": return string(0xF120)
		case "py", "pl", "PL", "t": return string(0xE769)
		case "msql", "mysql": return string(0xE229)
		case "pgsql", "sql": return string(0xF1C0)
		case "tcl": return string(0xE7C4)
		case "r", "R": return "R"
		case "gs": return "G"
		case "asm": return string(0xE79D)
		case "cl", "lisp": return string(0xE768)
		case "lua": return string(0xE620)
		case "moon": return string(0xF186)
		case "c", "C", "h", "H", "tcc", "M", "m": return string(0xE61E)
		case "c++", "cpp",  "h++", "hpp", "hxx", "ii", "cxx": return string(0xE61D)
		case "cc", "cs", "cp": return "#"
		case "cr": return string(0xE739)
		case "go": return string(0xE626)
		case "f", "for", "ftn": return "F"
		case "s", "S": return string(0xE79D)
		case "rs": return string(0xE7A8)
		case "sx": return "?"
		case "hi": return "I"
		case "hs", "lhs": return string(0xE61F)
		case "pyc": return string(0xE606)
		case "css": return string(0xE614)
		case "less": return string(0xE60B)
		case "sass", "scss": return string(0xE603)
		case "htm", "html", "jhtm", "mht", "eml": return string(0xE60E)
		case "mustache": return string(0xE60F)
		case "coffee", "java": return string(0xE61B)
		case "js", "jsm", "jsp": return string(0xE60C)
		case "php", "ctp": return string(0xE608)
		case "twig": return string(0xE61C)
		case "vb", "vba", "vbs": return "V"
		case "dockerignore": return string(0xE7B0)
		case "am", "in", "hin", "scan", "m4", "old", "out", "SKIP": return string(0xF0AD)
		case "diff", "patch": return string(0xF467)
    case "psd", "ps": return string(0xE7B8)
    case "ai": return string(0xE7B4)
		case "bmp", "tiff", "tif", "TIFF", "cdr", "gif", "ico", "jpeg", "JPG", "jpg", "nth", "png",
         "xpm", "epsf", "drw", "svg", "avi", "divx", "IFO", "m2v", "m4v", "mkv", "MOV",
         "mov", "mp4", "mpeg", "mpg", "ogm", "rmvb", "sample", "wmv", "VOB", "vob": return string(0xE60D)
		case "3g2", "3gp", "gp3", "webm", "gp4", "asf", "flv", "ts", "ogv", "f4v": return string(0xF10B)
		case "3ga", "S3M", "aac", "au", "dat", "dts", "fcm", "m4a", "mid", "midi", "mod", "mp3",
         "mp4a", "oga", "ogg", "opus", "s3m", "sid", "wma", "ape", "aiff", "cda", "flac", "alac",
         "pcm", "wav", "wv", "wvc": return string(0xF025)
		case "afm", "fon", "fnt", "pfb", "pfm", "ttf", "otf", "PFA", "pfa": return string(0xF031)
		case "7z", "a", "arj", "bz2", "cpio", "gz", "lrz", "lz", "lzma", "lzo", "rar", "s7z", "sz",
         "tar", "tgz", "xz", "zip", "zipx", "zoo", "zpaq", "zz", "Z", "z": return string(0xF187)
		case "apk", "deb", "rpm", "jad", "jar", "cab", "pak", "pk3", "vdf", "vpk", "bsp", "dmg": return string(0xF487)
		case "r00", "r01", "r02", "r03", "r04", "r05", "r06", "r07", "r08", "r09", "r10", "r100", "r101",
         "r102", "r103", "r104", "r105", "r106", "r107", "r108", "r109", "r11", "r110", "r111", "r112",
         "r113", "r114", "r115", "r116", "r12", "r13", "r14", "r15", "r16", "r17", "r18", "r19", "r20",
         "r21", "r22", "r25", "r26", "r27", "r28", "r29", "r30", "r31", "r32", "r33", "r34", "r35", "r36",
         "r37", "r38", "r39", "r40", "r41", "r42", "r43", "r44", "r45", "r46", "r47", "r48", "r49", "r50",
         "r51", "r52", "r53", "r54", "r55", "r56", "r57", "r58", "r59", "r60", "r61", "r62", "r63", "r64",
         "r65", "r66", "r67", "r68", "r69", "r70", "r71", "r72", "r73", "r74", "r75", "r76", "r77",
         "r78", "r79", "r80", "r81", "r82", "r83", "r84", "r85", "r86", "r87", "r88", "r89", "r90", "r91",
         "r92", "r93", "r94", "r95", "r96", "r97", "r98", "r99", "zx00", "zx01", "zx02", "zx03", "zx04",
         "zx05", "zx06", "zx07", "zx08", "zx09", "zx10", "zx11", "zx12", "zx13", "zx14", "zx15", "zx16",
         "zx17", "zx18", "zx19", "zx20", "zx21", "zx22", "zx25", "zx26", "zx27", "zx28", "zx29", "zx30",
         "zx31", "zx32", "zx33", "zx34", "zx35", "zx36", "zx37", "zx38", "zx39", "zx40", "zx41", "zx42",
         "zx43", "zx44", "zx45", "zx46", "zx47", "zx48", "zx49", "zx50", "zx51", "zx52", "zx53", "zx54",
         "zx55", "zx56", "zx57", "zx58", "zx59", "zx60", "zx61", "zx62", "zx63", "zx64", "zx65", "zx66",
         "zx67", "zx68", "zx69", "zx70", "zx71", "zx72", "zx73", "zx74", "zx75", "zx76", "zx77",
         "zx78", "zx79", "zx80", "zx81", "zx82", "zx83", "zx84", "zx85", "zx86", "zx87", "zx88", "zx89",
         "zx90", "zx91", "zx92", "zx93", "zx94", "zx95", "zx96", "zx97", "zx98", "zx99", "zx100", "zx101",
         "zx102", "zx103", "zx104", "zx105", "zx106", "zx107", "zx108", "zx109", "zx110", "zx111", "zx112",
         "zx113", "zx114", "zx115", "zx116", "z00", "z01", "z02", "z03", "z04", "z05", "z06", "z07", "z08",
         "z09", "z10", "z11", "z12", "z13", "z14", "z15", "z16", "z17", "z18", "z19", "z20", "z21", "z22",
         "z25", "z26", "z27", "z28", "z29", "z30", "z31", "z32", "z33", "z34", "z35", "z36", "z37", "z38",
         "z39", "z40", "z41", "z42", "z43", "z44", "z45", "z46", "z47", "z48", "z49", "z50", "z51", "z52",
         "z53", "z54", "z55", "z56", "z57", "z58", "z59", "z60", "z61", "z62", "z63", "z64", "z65", "z66",
         "z67", "z68", "z69", "z70", "z71", "z72", "z73", "z74", "z75", "z76", "z77", "z78", "z79",
         "z80", "z81", "z82", "z83", "z84", "z85", "z86", "z87", "z88", "z89", "z90", "z91", "z92", "z93",
         "z94", "z95", "z96", "z97", "z98", "z99", "z100", "z101", "z102", "z103", "z104", "z105", "z106",
         "z107", "z108", "z109", "z110", "z111", "z112", "z113", "z114", "z115", "z116", "part": return string(0xE601)
		case "iso", "bin", "nrg", "qcow", "sparseimage", "toast", "vcd", "vmdk": return string(0xF0A0)
		case "accdb", "accde", "accdr", "accdt", "db", "fmp12", "fp7", "localstorage", "mdb", "mde",
         "sqlite", "typelib", "nc": return string(0xF1C0)
		case "pacnew", "un", "orig", "BUP", "bak": return string(0xF0E2)
		case "swp", "swo", "tmp", "sassc": return "T"
		case "pid", "state": return string(0xF023)
		case "err", "error", "stderr", "deny": return string(0xF12A)
		case "dump", "stackdump", "zcompdump", "zwc": return string(0xF487)
		case "pcap", "cap", "dmp": return string(0xE765)
		case "DS_Store", "localized", "CFUserTextEncoding": return string(0xF179)
		case "allow": return string(0xF00C)
		case "service", "@.service", "socket", "swap", "device", "mount", "automount", "target", "path",
         "timer", "snapshot": return string(0xF109)
    case "md5": return "#"
		case "application", "cue", "description", "directory", "m3u", "m3u8", "properties", "sfv", "srt", "theme",
         "torrent", "urlview": return string(0xE60B)
		case "asc", "bfe", "enc", "gpg", "sign", "sig", "p12", "pem", "pgp": return string(0xF084)
		case "32x", "cdi", "fm2", "rom", "sav", "st", "a00", "a52", "A64", "a64", "a78", "adf", "atr", "gb", "gba",
         "gbc", "gel", "gg", "ggl", "j64", "nds", "nes", "sms": return string(0xF11B)
		case "pot": return "P"
    case "spl": return "S"
		case "pcb": return string(0xF493)
		case "mm": return string(0xF035)
		case "gbr", "scm", "Rproj": return string(0xF1FC)
		case "sis", "1p", "3p", "cnc", "def", "ex", "example", "feature", "ger", "map", "mf", "mfasl", "mi",
         "mtx", "pc", "pi", "plt", "pm", "rb", "rdf", "rst", "ru", "sch", "sty", "sug", "tdy", "tfm",
         "tfnt", "tg", "vcard", "vcf", "xln": return string(0xF10B)
		case "sc", "scala": return string(0xE737)
  }

  return default_icon

}
