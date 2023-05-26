interface Props {
  color?: string;
  width?: string;
  height?: string;
}

export default function ResendEmail({
  color = "#fff",
  width = "25px",
  height = "25px",
}: Props) {
  return (
    <svg
      style={{
        display: "inline-block",
      }}
      version="1.0"
      xmlns="http://www.w3.org/2000/svg"
      width={width}
      height={height}
      viewBox={`0 0 1024.000000 1024.000000`}
      preserveAspectRatio="xMidYMid meet"
    >
      <g
        transform="translate(0.000000,1024.000000) scale(0.100000,-0.100000)"
        fill={color}
        stroke="none"
      >
        <path
          d="M7359 9001 c-25 -10 -128 -61 -230 -114 -101 -52 -393 -201 -649
-330 -740 -373 -728 -366 -780 -471 -56 -113 -49 -226 20 -320 66 -89 131
-109 410 -126 183 -11 439 -35 445 -42 5 -5 -160 -481 -221 -638 -81 -211 -94
-254 -94 -315 0 -154 130 -287 280 -287 108 -1 234 73 267 157 7 17 42 107 78
200 102 264 126 314 152 315 39 0 553 -40 698 -55 177 -19 268 -19 331 0 163
48 256 197 224 360 -15 75 -44 142 -77 177 -18 19 -37 68 -68 175 -23 81 -90
285 -148 453 -58 168 -136 395 -173 504 -38 113 -81 220 -99 245 -33 51 -107
105 -163 121 -56 15 -152 11 -203 -9z m-5 -798 c23 -73 170 -498 200 -581 15
-41 26 -75 24 -77 -3 -2 -246 16 -396 31 -60 5 -63 7 -48 23 30 33 56 127 56
201 0 121 -47 206 -137 247 l-53 25 42 23 c24 12 97 52 163 88 66 35 123 65
127 66 3 0 13 -20 22 -46z"
        />
        <path
          d="M4605 7629 c-150 -11 -479 -54 -605 -80 -490 -99 -1006 -340 -1228
-573 -102 -106 -134 -170 -134 -260 0 -178 188 -316 337 -247 22 10 72 46 110
80 131 116 312 226 511 311 351 151 705 217 1244 235 169 5 202 3 355 -25 65
-12 76 -12 141 8 128 39 209 129 221 246 11 111 -67 236 -179 287 -52 23 -60
24 -353 25 -165 1 -354 -2 -420 -7z"
        />
        <path
          d="M4941 6649 c-451 -21 -760 -76 -1052 -185 -164 -61 -260 -111 -404
-207 -189 -128 -296 -251 -331 -380 -34 -126 41 -271 170 -327 86 -38 190 -7
246 75 87 127 345 279 614 362 270 83 690 123 1081 103 82 -5 95 -3 145 21 74
35 110 67 145 131 26 45 30 64 30 128 0 64 -4 83 -29 127 -48 87 -134 136
-269 153 -88 11 -101 11 -346 -1z"
        />
        <path
          d="M8130 5706 c-90 -20 -114 -33 -465 -251 -60 -37 -148 -90 -195 -118
-47 -27 -101 -60 -120 -72 -33 -21 -449 -277 -480 -295 -8 -5 -125 -76 -259
-157 -790 -479 -829 -501 -981 -568 -156 -70 -299 -105 -450 -112 -379 -17
-612 79 -1410 582 -85 54 -213 133 -285 175 -71 43 -249 151 -395 240 -146 90
-294 181 -330 203 -36 22 -114 71 -175 110 -295 188 -366 228 -451 253 -78 22
-151 14 -231 -27 -76 -39 -115 -81 -157 -169 l-31 -65 0 -1990 1 -1990 23 -58
c30 -76 115 -168 189 -204 l57 -28 3125 -3 c2264 -2 3139 0 3177 8 66 14 164
77 202 131 17 22 41 67 53 98 l23 56 3 1839 c1 1012 -1 1914 -6 2005 -7 147
-11 172 -32 216 -47 94 -133 162 -245 191 -60 16 -82 16 -155 0z m-150 -768
c0 -7 1 -733 2 -1614 l3 -1601 -35 27 c-19 16 -87 58 -150 96 -63 37 -182 110
-265 161 -82 52 -231 144 -330 205 -408 252 -1157 720 -1330 832 -207 134
-584 358 -635 377 -65 25 -164 22 -244 -6 -56 -19 -449 -253 -756 -448 -69
-44 -144 -92 -168 -106 -24 -14 -57 -35 -75 -46 -151 -97 -426 -267 -531 -330
-71 -42 -159 -97 -195 -120 -72 -47 -345 -215 -401 -247 -19 -11 -54 -33 -78
-49 -24 -16 -45 -29 -48 -29 -4 0 -267 -160 -397 -242 -23 -14 -44 -27 -47
-27 -3 -1 -5 708 -5 1574 0 866 1 1575 1 1575 1 0 39 -25 85 -57 71 -47 254
-161 439 -273 25 -15 81 -49 125 -77 102 -63 195 -120 265 -163 30 -18 120
-74 200 -123 669 -415 831 -488 1195 -542 130 -19 755 -23 902 -5 368 44 536
111 1043 413 102 61 237 140 300 177 63 37 243 145 400 240 157 95 377 228
490 296 113 68 210 128 215 133 14 14 25 13 25 -1z m-2821 -2135 c25 -15 118
-74 206 -131 88 -58 254 -163 370 -235 115 -72 277 -173 358 -224 81 -51 150
-93 152 -93 2 0 45 -26 94 -57 128 -84 287 -184 440 -279 l135 -84 -1790 0
c-985 0 -1794 3 -1797 6 -4 4 17 21 46 39 29 17 99 61 157 97 473 296 544 340
656 408 62 38 154 95 205 128 93 60 493 314 637 405 42 26 78 47 81 47 2 0 25
-12 50 -27z"
        />
        <path
          d="M27 103 c-25 -4 -27 -8 -27 -54 0 -48 1 -49 30 -49 30 0 30 1 30 55
0 30 -1 54 -2 54 -2 -1 -15 -4 -31 -6z"
        />
      </g>
    </svg>
  );
}
