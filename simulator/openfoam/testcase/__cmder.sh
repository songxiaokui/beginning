set -e
# clean old time series data
foamListTimes -rm -noZero

# mesh transform *.unv
ideasUnvToFoam *.unv

BOUNDARY_FILE="constant/polyMesh/boundary"
P_FILE="0/p"
U_FILE="0/U"

# add others scriptes
perl -0777 -i.bak -pe '
s/(\b[\w-]*face_top[\w-]*\s*\{\s*type\s+)patch(\s*;)/${1}empty$2/gis;
s/(\b[\w-]*face[\w-]*\s*\{\s*type\s+)patch(\s*;)/${1}empty$2/gis;

s/(\b[\w-]*wall[\w-]*\s*\{\s*type\s+)patch(\s*;)/${1}wall$2/gis;
s/(\b[\w-]*symmetry[\w-]*\s*\{\s*type\s+)patch(\s*;)/${1}symmetry$2/gis;

s/(\b[\w-]*outlet[\w-]*\s*\{\s*type\s+)patch(\s*;)/${1}outlet$2/gis;
s/(\b[\w-]*inlet[\w-]*\s*\{\s*type\s+)patch(\s*;)/${1}inlet$2/gis;
' "${BOUNDARY_FILE}"

# solve the problem that all walls are not checked
# patch defaultFaces
if grep -Eq '^[[:space:]]*defaultFaces([[:space:]]*\{)?[[:space:]]*$' "$BOUNDARY_FILE"; then
    echo "[INFO] defaultFaces found in $BOUNDARY_FILE"

    # boundary: defaultFaces type empty -> wall
    perl -0777 -i -pe '
s/(\bdefaultFaces\s*\{\s*type\s+)empty(\s*;)/${1}wall$2/gis;
' "$BOUNDARY_FILE"

    echo "[INFO] defaultFaces type changed to wall"


    # patch 0/p
    if ! grep -Eq '^[[:space:]]*defaultFaces([[:space:]]*\{)?[[:space:]]*$' "$P_FILE"; then
        echo "[INFO] patch defaultFaces into $P_FILE"

        if perl -0777 -i -pe '
s/\n\}[[:space:]]*$/\n    defaultFaces\n    {\n        type    zeroGradient;\n\n    }\n}\n/s;
' "$P_FILE"; then
            echo "[INFO] defaultFaces patched into $P_FILE"
        else
            echo "[ERROR] failed to patch defaultFaces into $P_FILE"
            exit 1
        fi
    else
        echo "[INFO] defaultFaces already exists in $P_FILE, skip"
    fi


    # patch 0/U
    if ! grep -Eq '^[[:space:]]*defaultFaces([[:space:]]*\{)?[[:space:]]*$' "$U_FILE"; then
        echo "[INFO] patch defaultFaces into $U_FILE"

        if perl -0777 -i -pe '
s/\n\}[[:space:]]*$/\n    defaultFaces\n    {\n        type    noSlip;\n\n    }\n}\n/s;
' "$U_FILE"; then
            echo "[INFO] defaultFaces patched into $U_FILE"
        else
            echo "[ERROR] failed to patch defaultFaces into $U_FILE"
            exit 1
        fi
    else
        echo "[INFO] defaultFaces already exists in $U_FILE, skip"
    fi

else
    echo "[INFO] defaultFaces not found in $BOUNDARY_FILE, skip patch"
fi

# solver
# to change your solver , you should change it in 
# support: icoFoam、simpleFoam
simpleFoam

# create flag file
touch results.foam
