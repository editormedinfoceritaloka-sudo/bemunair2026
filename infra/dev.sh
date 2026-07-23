GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' 

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$SCRIPT_DIR"


function info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

function warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

function error() {
    echo -e "${RED}[ERROR]${NC} $1"
}


function usage() {
    echo "Usage: ./dev.sh [command]"
    echo ""
    echo "Commands:"
    echo "  up             Start containers with watch mode"
    echo "  down           Stop and remove containers"
    echo "  build          Build or rebuild services without cache"
    echo "  logs [service] Tail logs (all or specific service)"
    echo "  reset          Stop, remove volumes, and start again"
    echo "  ps             List container status"
    echo ""
}

case "$1" in
    up)
        info "Starting containers with watch mode..."
        docker compose watch
        ;;
    down)
        info "Stopping containers..."
        docker compose down
        ;;
    build)
        info "Building containers without cache..."
        docker compose build --no-cache
        ;;
    logs)
        info "Tailing logs..."
        docker compose logs -f "$2"
        ;;
    reset)
        warn "Resetting environment (removing volumes)..."
        docker compose down -v
        info "Starting fresh..."
        docker compose up -d
        info "Containers started in background. Use './dev.sh logs' to see logs."
        ;;
    ps)
        info "Container status:"
        docker compose ps
        ;;
    *)
        usage
        ;;
esac
