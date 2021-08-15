package base64

import (
	"encoding/base64"
	"math/rand"
)

var (
	valueString             = `THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.`
	valueBytes              = []byte(valueString)
	stdBase64ValueString    = base64.StdEncoding.EncodeToString(valueBytes)
	stdBase64ValueBytes     = []byte(stdBase64ValueString)
	rawStdBase64ValueString = base64.RawStdEncoding.EncodeToString(valueBytes)
	rawStdBase64ValueBytes  = []byte(rawStdBase64ValueString)
	urlBase64ValueString    = base64.URLEncoding.EncodeToString(valueBytes)
	urlBase64ValueBytes     = []byte(urlBase64ValueString)
	rawURLBase64ValueString = base64.URLEncoding.EncodeToString(valueBytes)
	rawURLBase64ValueBytes  = []byte(rawURLBase64ValueString)

	byteResult   []byte
	stringResult string
	resultN      int
)

func generateRandomBytes(length int) []byte {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		for result[i] == 0 || result[i] == 10 || result[i] == 13 {
			result[i] = byte(rand.Uint32())
		}
	}
	return result
}
